// package main is a module with raspberry pi board component.
package main

import (
	"context"

	"github.com/viam-modules/gps/nmea"
	"github.com/viam-modules/gps/rtk"
	dualgps "github.com/viam-modules/gps/rtk-dual-gps"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("gps"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	module, err := module.NewModuleFromArgs(ctx)
	if err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, movementsensor.API, rtk.ModelSerial); err != nil {
		return err
	}

	/* if err = module.AddModelFromRegistry(ctx, movementsensor.API, rtk.ModelPmtk); err != nil {
		return err
	} */

	if err = module.AddModelFromRegistry(ctx, movementsensor.API, nmea.Model); err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, movementsensor.API, dualgps.Model); err != nil {
		return err
	}

	err = module.Start(ctx)
	defer module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
