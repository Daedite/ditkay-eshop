package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductRepository interface {
	CreateProduct(product entity.Product) (models.Product, error)
	ReadProduct(id string) (models.Product, error)
	UpdateProduct(product entity.Product) (models.Product, error)
	DeleteProduct(product entity.Product) (bool, error)
	ReadProducts() ([]gorm.Product, error)
}
