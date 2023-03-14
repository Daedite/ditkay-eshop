package eshop_api

type ConfigManager interface {
	Load() (ServiceConfiguration, error)
}
