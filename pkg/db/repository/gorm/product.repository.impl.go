package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	GormDB *gorm.DB
}

func NewProductRepositoryImpl(gormDB *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		GormDB: gormDB,
	}
}

var _ repository.ProductRepository = &ProductRepositoryImpl{}

func (p ProductRepositoryImpl) CreateProduct(product entity.Product) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) ReadProduct(id string) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) UpdateProduct(product entity.Product) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) DeleteProduct(product entity.Product) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) ReadProducts() ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}
