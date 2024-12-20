// Package nmea implements an NMEA gps.
package nmea

import (
	"context"

	"github.com/viam-modules/gps/gpsutils"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

// NewSerialGPSNMEA creates a component that communicates over a serial port.
func NewSerialGPSNMEA(ctx context.Context, name resource.Name, conf *Config, logger logging.Logger) (movementsensor.MovementSensor, error) {
	dev, err := gpsutils.NewSerialDataReader(ctx, conf.SerialConfig, logger)
	if err != nil {
		return nil, err
	}

	return newNMEAMovementSensor(ctx, name, dev, logger)
}
