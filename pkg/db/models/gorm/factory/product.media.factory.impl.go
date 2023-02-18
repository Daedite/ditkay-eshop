package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaFactoryImpl struct {
}

func (p ProductMediaFactoryImpl) CreateProductMediaFactory(productMedia entity.ProductMedia) models.ProductMedia {
	return &gorm.ProductMedia{
		Id:        productMedia.Id,
		MediaId:   productMedia.MediaId,
		MediaType: productMedia.MediaType,
		ProductId: productMedia.ProductId,
	}
}

func NewProductMediaFactoryImpl() *ProductMediaFactoryImpl {
	return &ProductMediaFactoryImpl{}
}

var _ factory.ProductMediaFactory = &ProductMediaFactoryImpl{}
