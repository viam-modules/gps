package rtk

import (
	"context"
	"errors"
	"math"
	"testing"

	geo "github.com/kellydunn/golang-geo"
	"github.com/viam-modules/gps/gpsutils"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/test"
)

type mockDataReader struct{}

func (d *mockDataReader) Messages() chan string {
	return nil
}

func (d *mockDataReader) Close() error {
	return nil
}

func TestValidateSerialRTK(t *testing.T) {
	path := "path"
	t.Run("valid config", func(t *testing.T) {
		cfg := SerialConfig{
			NtripURL:             "http//fakeurl",
			NtripConnectAttempts: 10,
			NtripPass:            "somepass",
			NtripUser:            "someuser",
			NtripMountpoint:      "NYC",
			SerialPath:           path,
			SerialBaudRate:       115200,
		}
		_, err := cfg.Validate(path)
		test.That(t, err, test.ShouldBeNil)
	})

	t.Run("invalid url config", func(t *testing.T) {
		cfg := SerialConfig{
			NtripURL:             "",
			NtripConnectAttempts: 10,
			NtripPass:            "somepass",
			NtripUser:            "someuser",
			NtripMountpoint:      "NYC",
			SerialPath:           path,
			SerialBaudRate:       115200,
		}

		_, err := cfg.Validate(path)
		test.That(t, err, test.ShouldBeError,
			resource.NewConfigValidationFieldRequiredError(path, "ntrip_url"))
	})

	t.Run("invalid path config", func(t *testing.T) {
		cfg := SerialConfig{
			NtripURL:             "http//fakeurl",
			NtripConnectAttempts: 10,
			NtripPass:            "somepass",
			NtripUser:            "someuser",
			NtripMountpoint:      "NYC",
			SerialPath:           "",
			SerialBaudRate:       115200,
		}

		_, err := cfg.Validate(path)
		test.That(t, err, test.ShouldBeError,
			resource.NewConfigValidationFieldRequiredError(path, "serial_path"))
	})
}

// This sets the position to 12°34.5678' N, 123°45.6789' W, at time 12:34:56.78 UTC.
const setPositionSentence = "$GPGLL,1234.5678,N,12345.6789,W,123456.78,A,D*7F"

// initializePosition sets the position in the cached data and returns the point it is set to.
func initializePosition(cachedData *gpsutils.CachedData) *geo.Point {
	cachedData.ParseAndUpdate(setPositionSentence)
	return geo.NewPoint(12.576130000000001, -123.76131500000001)
}

func TestPosition(t *testing.T) {
	// WITH LAST ERROR

	// If there is last error and no last position, return NaN
	t.Run("position with last error and no last position", func(t *testing.T) {
		g := &gpsrtk{
			err: movementsensor.NewLastError(1, 1),
		}
		g.err.Set(errors.New("last error test"))

		pos, alt, err := g.Position(context.Background(), nil)
		test.That(t, movementsensor.IsPositionNaN(pos), test.ShouldBeTrue)
		test.That(t, math.IsNaN(alt), test.ShouldBeTrue)
		test.That(t, err, test.ShouldBeError, "last error test")
	})

	// If there is last error and last position, return last position
	t.Run("position with last error and last position", func(t *testing.T) {
		g := &gpsrtk{
			err:        movementsensor.NewLastError(1, 1),
			cachedData: gpsutils.NewCachedData(&mockDataReader{}, logging.NewTestLogger(t)),
		}
		initializePosition(g.cachedData)
		g.err.Set(errors.New("last position"))

		pos, alt, err := g.Position(context.Background(), nil)
		test.That(t, math.IsNaN(pos.Lat()), test.ShouldBeTrue)
		test.That(t, math.IsNaN(pos.Lng()), test.ShouldBeTrue)
		test.That(t, math.IsNaN(alt), test.ShouldBeTrue)
		test.That(t, err, test.ShouldBeError, "last position")
	})

	// NO LAST ERROR, but with cachedData ERROR

	// If there is no last error, invalid current position and no last position, return NaN
	t.Run("invalid position with invalid last position, with position error", func(t *testing.T) {
		g := &gpsrtk{
			err:        movementsensor.NewLastError(1, 1),
			cachedData: gpsutils.NewCachedData(&mockDataReader{}, logging.NewTestLogger(t)),
		}

		pos, alt, err := g.Position(context.Background(), nil)
		test.That(t, movementsensor.IsPositionNaN(pos), test.ShouldBeTrue)
		test.That(t, math.IsNaN(alt), test.ShouldBeTrue)
		test.That(t, err, test.ShouldBeError, "nil gps location, check nmea message parsing")
	})

	// If there is no last error, invalid current position and valid last position, return last position
	t.Run("invalid position with valid last position, with position error", func(t *testing.T) {
		g := &gpsrtk{
			err:        movementsensor.NewLastError(1, 1),
			cachedData: gpsutils.NewCachedData(&mockDataReader{}, logging.NewTestLogger(t)),
		}
		expectedPos := initializePosition(g.cachedData)

		// This is an invalid command, containing too many periods and colons.
		invalidPosition := "$GPGLL,87.65.4321,N,987.65.4321,W,12:34:56.78,A,D*7F"
		g.cachedData.ParseAndUpdate(invalidPosition)

		pos, _, err := g.Position(context.Background(), nil)
		test.That(t, pos, test.ShouldResemble, expectedPos)
		test.That(t, err, test.ShouldBeNil)
	})

	// NO ERRORS

	// Invalid current position from NMEA message, return last known position
	t.Run("invalid position with valid last position, no error", func(t *testing.T) {
		g := &gpsrtk{
			err:        movementsensor.NewLastError(1, 1),
			cachedData: gpsutils.NewCachedData(&mockDataReader{}, logging.NewTestLogger(t)),
		}
		expectedPos := initializePosition(g.cachedData)

		// NMEA sentence with invalid position, Fix quality is 0
		nmeaSentenceInvalid := "$GPGGA,172814.0,123.123,N,234.234,W,0,6,1.2,18.893,M,-25.669,M,2.0,0031*4F"
		g.cachedData.ParseAndUpdate(nmeaSentenceInvalid)

		pos, _, err := g.Position(context.Background(), nil)
		test.That(t, pos, test.ShouldResemble, expectedPos)
		test.That(t, err, test.ShouldBeNil)
	})

	// Valid current position, should return current position
	t.Run("valid position, no error", func(t *testing.T) {
		g := &gpsrtk{
			err:        movementsensor.NewLastError(1, 1),
			cachedData: gpsutils.NewCachedData(&mockDataReader{}, logging.NewTestLogger(t)),
		}

		// Valid NMEA sentence
		nmeaSentenceValid := "$GPGGA,172814.0,3723.46587704,N,12202.26957864,W,2,6,1.2,18.893,M,-25.669,M,2.0,0031*4F"
		g.cachedData.ParseAndUpdate(nmeaSentenceValid)

		pos, _, err := g.Position(context.Background(), nil)

		expectedLat := (37 + 23.46587704/60)
		expectedLng := -(122 + 2.26957864/60)

		test.That(t, pos.Lat(), test.ShouldAlmostEqual, expectedLat)
		test.That(t, pos.Lng(), test.ShouldAlmostEqual, expectedLng)
		test.That(t, err, test.ShouldBeNil)
	})
}
