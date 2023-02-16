package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"gorm.io/gorm"
)

type ProductMediaRepository struct {
	GormDB *gorm.DB
}

func NewProductMediaRepository(gormdb *gorm.DB) *ProductMediaRepository {
	return &ProductMediaRepository{
		GormDB: gormdb,
	}
}

var _ repository.ProductMediaRepository = &ProductMediaRepository{}

func (p ProductMediaRepository) CreateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductMediaRepository) ReadProductMedia(id string) (models.ProductMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductMediaRepository) UpdateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductMediaRepository) DeleteProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductMediaRepository) ReadProductMedias() ([]models.ProductMedia, error) {
	//TODO implement me
	panic("implement me")
}
