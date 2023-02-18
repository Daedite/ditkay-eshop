package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	gormModel "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaRepository interface {
	CreateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	ReadProductMedia(id int) (models.ProductMedia, error)
	UpdateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	DeleteProductMedia(productMedia entity.ProductMedia) (bool, error)
	ReadProductMedias() ([]gormModel.ProductMedia, error)
}
