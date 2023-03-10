package factory

import (
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeFactory interface {
	CreateMediaType(body spec.MediaType) (*entity.MediaType, error)
}
type MediaTypeFactoryImpl struct {
}

func NewMediaTypeFactoryImpl() *MediaTypeFactoryImpl {
	return &MediaTypeFactoryImpl{}
}

func (m MediaTypeFactoryImpl) CreateMediaType(body spec.MediaType) (*entity.MediaType, error) {

	return &entity.MediaType{
		Id:          *body.Id,
		Name:        *body.Name,
		Description: *body.Description,
	}, nil
}

var _ MediaTypeFactory = &MediaTypeFactoryImpl{}
