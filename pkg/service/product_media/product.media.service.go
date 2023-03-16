package product_media

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaService interface {
	CreateProductMedia(productMedia entity.ProductMedia) (*entity.ProductMedia, error)
	ReadProductMedia(id string) (*entity.ProductMedia, error)
	UpdateProductMedia(productMedia entity.ProductMedia) (*entity.ProductMedia, error)
	DeleteProductMedia(productMedia entity.ProductMedia) (bool, error)
	ReadProductMedias() (*[]entity.ProductMedia, error)
	ReadProductMediasByProductId(productId string) ([]entity.Media, error)
	RemoveMedias(mediaId string) (bool, error)
}
