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

type ProductMediaRepositoryImpl struct {
	GormDB  *gorm.DB
	Factory factory.ProductMediaFactory
}

func NewProductMediaRepositoryImpl(gormdb *gorm.DB, factory factory.ProductMediaFactory) *ProductMediaRepositoryImpl {
	return &ProductMediaRepositoryImpl{
		GormDB:  gormdb,
		Factory: factory,
	}
}

var _ repository.ProductMediaRepository = &ProductMediaRepositoryImpl{}

func (p ProductMediaRepositoryImpl) CreateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error) {
	var gormProductMedia *gormModel.ProductMedia = p.Factory.CreateProductMediaFactory(productMedia).(*gormModel.ProductMedia)
	if err := p.GormDB.Create(&gormProductMedia).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to create Product media id: %d", productMedia.Id))
		return nil, err
	}
	return gormProductMedia, nil
}

func (p ProductMediaRepositoryImpl) ReadProductMedia(id int) (models.ProductMedia, error) {
	gormProductMedia := gormModel.ProductMedia{}
	if err := p.GormDB.Where("id = ?", id).First(&gormProductMedia).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to get Product media id: %d", id))
		return nil, err
	}
	return gormProductMedia, nil
}

func (p ProductMediaRepositoryImpl) UpdateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error) {
	var gormProduct *gormModel.ProductMedia = p.Factory.CreateProductMediaFactory(productMedia).(*gormModel.ProductMedia)
	if err := p.GormDB.Model(&gormModel.Product{}).Where("id = ? ", productMedia.Id).Updates(productMedia).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update product media id: %d", productMedia.Id))
		return nil, err
	}
	return gormProduct, nil
}

func (p ProductMediaRepositoryImpl) DeleteProductMedia(productMedia entity.ProductMedia) (bool, error) {
	gormProductMedia := &gormModel.ProductMedia{}
	if err := p.GormDB.Where("id = ?", productMedia.Id).Delete(&gormProductMedia).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to delete product media id: %d", productMedia.Id))
		return false, err
	}
	return true, nil
}

func (p ProductMediaRepositoryImpl) ReadProductMedias() ([]gormModel.ProductMedia, error) {
	gormProduct := []gormModel.ProductMedia{}
	if err := p.GormDB.Find(&gormProduct).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads product medias"))
		return nil, err
	}
	return gormProduct, nil
}
