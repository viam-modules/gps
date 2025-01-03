//go:build linux

// Package nmea implements a GPS NMEA component.
package nmea

import (
	"context"

	"github.com/viam-modules/gps/gpsutils"
	"go.viam.com/rdk/components/board/genericlinux/buses"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

// NewPmtkI2CGPSNMEA implements a gps that communicates over i2c.
func NewPmtkI2CGPSNMEA(
	ctx context.Context,
	deps resource.Dependencies,
	name resource.Name,
	conf *Config,
	logger logging.Logger,
) (movementsensor.MovementSensor, error) {
	// The nil on this next line means "use a real I2C bus, because we're not going to pass in a
	// mock one."
	return MakePmtkI2cGpsNmea(ctx, deps, name, conf, logger, nil)
}

// MakePmtkI2cGpsNmea is only split out for ease of testing: you can pass in your own mock I2C bus,
// or pass in nil to have it create a real one. It is public so it can also be called from within
// the gpsrtkpmtk package.
func MakePmtkI2cGpsNmea(
	ctx context.Context,
	deps resource.Dependencies,
	name resource.Name,
	conf *Config,
	logger logging.Logger,
	i2cBus buses.I2C,
) (movementsensor.MovementSensor, error) {
	dev, err := gpsutils.NewI2cDataReader(ctx, *conf.I2CConfig, i2cBus, logger)
	if err != nil {
		return nil, err
	}

	return newNMEAMovementSensor(ctx, name, dev, logger)
}
