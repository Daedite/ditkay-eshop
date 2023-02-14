package env

import (
	eshop_api "github.com/ESPOIR-DITE/ditkay-eshop/pkg/config/eshop-api"
	"github.com/kelseyhightower/envconfig"
)

type DitkayEshopApiConfigManagerImpl struct{}

func NewDitkayEshopApiConfigManagerImpl() *DitkayEshopApiConfigManagerImpl {
	return &DitkayEshopApiConfigManagerImpl{}
}

func (e DitkayEshopApiConfigManagerImpl) Load() (config eshop_api.ServiceConfiguration, err error) {
	config = new(EshopEnvServiceConfiguration)
	if err = envconfig.Process("", config); err != nil {
		return
	}
	return
}

var _ eshop_api.ConfigManager = &DitkayEshopApiConfigManagerImpl{}
