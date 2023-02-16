package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaFactory interface {
	CreateProductMediaFactory(productMedia entity.ProductMedia) models.ProductMedia
}
