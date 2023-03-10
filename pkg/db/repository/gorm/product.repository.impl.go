package gorm

import (
	"fmt"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	gormModel "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	GormDB  *gorm.DB
	Factory factory.ProductFactory
}

func NewProductRepositoryImpl(gormDB *gorm.DB, factory factory.ProductFactory) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		GormDB:  gormDB,
		Factory: factory,
	}
}

var _ repository.ProductRepository = &ProductRepositoryImpl{}

func (p ProductRepositoryImpl) CreateProduct(product entity.Product) (models.Product, error) {
	var gormProduct *gormModel.Product = p.Factory.CreateProductFactory(product).(*gormModel.Product)
	if err := p.GormDB.Create(&gormProduct).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to create Product id: %d", product.Id))
		return nil, err
	}
	return gormProduct, nil
}

func (p ProductRepositoryImpl) ReadProduct(id string) (models.Product, error) {
	gormProduct := gormModel.Product{}
	if err := p.GormDB.Where("id = ?", id).First(&gormProduct).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to get Product id: %d", id))
		return nil, err
	}
	return gormProduct, nil
}

func (p ProductRepositoryImpl) UpdateProduct(product entity.Product) (models.Product, error) {
	var gormProduct *gormModel.Product = p.Factory.CreateProductFactory(product).(*gormModel.Product)
	if err := p.GormDB.Model(&gormModel.Product{}).Where("id = ? ", product.Id).Updates(product).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update product id: %d", product.Id))
		return nil, err
	}
	return gormProduct, nil
}

func (p ProductRepositoryImpl) DeleteProduct(product entity.Product) (bool, error) {
	gormProduct := &gormModel.Product{}
	if err := p.GormDB.Where("id = ?", product.Id).Delete(&gormProduct).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to delete product id: %d", product.Id))
		return false, err
	}
	return true, nil
}

func (p ProductRepositoryImpl) ReadProducts() ([]gormModel.Product, error) {
	gormProduct := []gormModel.Product{}
	if err := p.GormDB.Find(&gormProduct).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads products"))
		return nil, err
	}
	return gormProduct, nil
}
