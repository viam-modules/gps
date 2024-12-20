// Package nmea implements an NMEA gps.
package nmea

import (
	"context"

	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"github.com/viam-modules/gps/gpsutils"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

// MovementSensor allows the use of any MovementSensor chip via a DataReader.
type MovementSensor struct {
	resource.Named
	resource.AlwaysRebuild
	logger     logging.Logger
	cachedData *gpsutils.CachedData
}

// newNMEAMovementSensor creates a new movement sensor.
func newNMEAMovementSensor(
	_ context.Context, name resource.Name, dev gpsutils.DataReader, logger logging.Logger,
) (movementsensor.MovementSensor, error) {
	g := &MovementSensor{
		Named:      name.AsNamed(),
		logger:     logger,
		cachedData: gpsutils.NewCachedData(dev, logger),
	}

	return g, nil
}

// Position returns the position and altitide of the sensor, or an error.
func (g *MovementSensor) Position(
	ctx context.Context, extra map[string]interface{},
) (*geo.Point, float64, error) {
	return g.cachedData.Position(ctx, extra)
}

// Accuracy returns the accuracy map, hDOP, vDOP, Fixquality and compass heading error.
func (g *MovementSensor) Accuracy(
	ctx context.Context, extra map[string]interface{},
) (*movementsensor.Accuracy, error) {
	return g.cachedData.Accuracy(ctx, extra)
}

// LinearVelocity returns the sensor's linear velocity. It requires having a compass heading, so we
// know which direction our speed is in. We assume all of this speed is horizontal, and not in
// gaining/losing altitude.
func (g *MovementSensor) LinearVelocity(
	ctx context.Context, extra map[string]interface{},
) (r3.Vector, error) {
	return g.cachedData.LinearVelocity(ctx, extra)
}

// LinearAcceleration returns the sensor's linear acceleration.
func (g *MovementSensor) LinearAcceleration(
	ctx context.Context, extra map[string]interface{},
) (r3.Vector, error) {
	return g.cachedData.LinearAcceleration(ctx, extra)
}

// AngularVelocity returns the sensor's angular velocity.
func (g *MovementSensor) AngularVelocity(
	ctx context.Context, extra map[string]interface{},
) (spatialmath.AngularVelocity, error) {
	return g.cachedData.AngularVelocity(ctx, extra)
}

// Orientation returns the sensor's orientation.
func (g *MovementSensor) Orientation(
	ctx context.Context, extra map[string]interface{},
) (spatialmath.Orientation, error) {
	return g.cachedData.Orientation(ctx, extra)
}

// CompassHeading returns the heading, from the range 0->360.
func (g *MovementSensor) CompassHeading(
	ctx context.Context, extra map[string]interface{},
) (float64, error) {
	return g.cachedData.CompassHeading(ctx, extra)
}

// Readings will use return all of the MovementSensor Readings.
func (g *MovementSensor) Readings(
	ctx context.Context, extra map[string]interface{},
) (map[string]interface{}, error) {
	readings, err := movementsensor.DefaultAPIReadings(ctx, g, extra)
	if err != nil {
		return nil, err
	}

	commonReadings := g.cachedData.GetCommonReadings(ctx)

	readings["fix"] = commonReadings.FixValue
	readings["satellites_in_use"] = commonReadings.SatsInUse

	return readings, nil
}

// Properties returns what movement sensor capabilities we have.
func (g *MovementSensor) Properties(
	ctx context.Context, extra map[string]interface{},
) (*movementsensor.Properties, error) {
	return g.cachedData.Properties(ctx, extra)
}

// Close shuts down the MovementSensor.
func (g *MovementSensor) Close(ctx context.Context) error {
	g.logger.CDebug(ctx, "Closing MovementSensor")
	// In some of the unit tests, the cachedData is nil. Only close it if it's not.
	if g.cachedData != nil {
		return g.cachedData.Close(ctx)
	}
	return nil
}
