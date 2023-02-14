package controller

import (
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
)

type DitKayEshopApiController struct {
}

func NewDitKayEshopApiController() *DitKayEshopApiController {
	return &DitKayEshopApiController{}
}

var _ spec.ServerInterface = DitKayEshopApiController{}
