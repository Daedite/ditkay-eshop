package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductFactory interface {
	CreateProductFactory(product entity.Product) models.Product
}
