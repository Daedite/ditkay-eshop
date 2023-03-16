package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	gormModel "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaRepository interface {
	CreateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	ReadProductMedia(id string) (models.ProductMedia, error)
	UpdateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	DeleteProductMedia(productMedia entity.ProductMedia) (bool, error)
	ReadProductMedias() ([]gormModel.ProductMedia, error)
	ReadProductMediasByProductId(productId string) ([]gormModel.ProductMedia, error)
	ReadProductMediaWithImageId(imageId string) (models.ProductMedia, error)
}
