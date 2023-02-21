package main

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config/eshop-api/env"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/wire"
	"os"
)

const version = "0.0.1"

func main() {
	logger.Log.Infof("starting Ditkay eshop Api v: %s", version)
	os.Exit(wire.Run(env.NewDitkayEshopApiConfigManagerImpl()))
}
