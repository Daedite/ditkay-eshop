package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductFactoryImpl struct {
}

func (p ProductFactoryImpl) CreateProductFactory(product entity.Product) models.Product {
	return &gorm.Product{
		Id:          product.Id,
		Name:        product.Name,
		SellPrice:   product.SellPrice,
		BuyPrice:    product.BuyPrice,
		Quantity:    product.Quantity,
		Description: product.Description,
	}
}

func NewProductFactoryImpl() *ProductFactoryImpl {
	return &ProductFactoryImpl{}
}

var _ factory.ProductFactory = &ProductFactoryImpl{}
