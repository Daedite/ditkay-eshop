package controller

import (
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media_type"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product_media"
)

type DitKayEshopApiController struct {
	MediaService        media.MediaService
	MediaTypeService    media_type.MediaTypeService
	ProductService      product.ProductService
	ProductMediaService product_media.ProductMediaService
	MediaFactory        factory.MediaFactory
	MediaTypeFactory    factory.MediaTypeFactory
	ProductFactory      factory.ProductFactory
	ProductMediaFactory factory.ProductMediaFactory
}

func NewDitKayEshopApiController(
	mediaService media.MediaService,
	mediaTypeService media_type.MediaTypeService,
	productService product.ProductService,
	productMediaService product_media.ProductMediaService,
	mediaFactory factory.MediaFactory,
	mediaTypeFactory factory.MediaTypeFactory,
	productFactory factory.ProductFactory,
	productMediaFactory factory.ProductMediaFactory,
) *DitKayEshopApiController {
	return &DitKayEshopApiController{
		MediaService:        mediaService,
		MediaTypeService:    mediaTypeService,
		ProductService:      productService,
		ProductMediaService: productMediaService,
		MediaFactory:        mediaFactory,
		MediaTypeFactory:    mediaTypeFactory,
		ProductFactory:      productFactory,
		ProductMediaFactory: productMediaFactory,
	}
}

var _ spec.ServerInterface = DitKayEshopApiController{}
