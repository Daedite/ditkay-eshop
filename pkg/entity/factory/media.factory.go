package factory

import (
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaFactory interface {
	CreateMedia(body spec.Media) (*entity.Media, error)
}
type MediaFactoryImpl struct {
}

func NewMediaFactoryImpl() *MediaFactoryImpl {
	return &MediaFactoryImpl{}
}

func (m MediaFactoryImpl) CreateMedia(body spec.Media) (*entity.Media, error) {

	return &entity.Media{
		Id:          *body.Id,
		Image:       *body.Image,
		Description: *body.Description,
	}, nil
}

var _ MediaFactory = &MediaFactoryImpl{}
