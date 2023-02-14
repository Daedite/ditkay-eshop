package wire

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"
	controller2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/http/controller"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/http/server"
	"gorm.io/gorm"
)

func wire(config config.ServiceConfiguration, gorm *gorm.DB) server.HttpServer {
	controller := controller2.NewDitKayEshopApiController()
	return server.NewHttpServerImpl(config.AppConfig(), controller)
}
