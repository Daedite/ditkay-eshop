package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaRepository interface {
	CreateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	ReadProductMedia(id string) (models.ProductMedia, error)
	UpdateProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	DeleteProductMedia(productMedia entity.ProductMedia) (models.ProductMedia, error)
	ReadProductMedias() ([]models.ProductMedia, error)
}
