package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaFactory struct {
}

func (p ProductMediaFactory) CreateProductMediaFactory(productMedia entity.ProductMedia) models.ProductMedia {
	return &gorm.ProductMedia{
		Id:        productMedia.Id,
		MediaId:   productMedia.MediaId,
		MediaType: productMedia.MediaType,
		ProductId: productMedia.ProductId,
	}
}

func NewProductMediaFactory() *ProductMediaFactory {
	return &ProductMediaFactory{}
}

var _ factory.ProductMediaFactory = &ProductMediaFactory{}
