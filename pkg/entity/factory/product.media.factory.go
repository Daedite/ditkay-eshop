package factory

import (
	"errors"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type ProductMediaFactory interface {
	CreateProductMedia(body spec.ProductMedia) (*entity.ProductMedia, error)
}
type ProductMediaFactoryImpl struct {
}

func NewProductMediaFactoryImpl() *ProductMediaFactoryImpl {
	return &ProductMediaFactoryImpl{}
}

var _ ProductMediaFactory = &ProductMediaFactoryImpl{}

func (m ProductMediaFactoryImpl) CreateProductMedia(body spec.ProductMedia) (*entity.ProductMedia, error) {
	if body.Id == nil {
		return nil, errors.New("missing required")
	}
	return &entity.ProductMedia{
		Id:        *body.Id,
		MediaId:   *body.MediaId,
		MediaType: *body.MediaType,
		ProductId: *body.ProductId,
	}, nil
}
