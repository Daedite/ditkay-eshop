package product

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductService interface {
	CreateProduct(product entity.Product) (*entity.Product, error)
	ReadProduct(id int) (*entity.Product, error)
	UpdateProduct(product entity.Product) (*entity.Product, error)
	DeleteProduct(product entity.Product) (bool, error)
	ReadProducts() (*[]entity.Product, error)
}
