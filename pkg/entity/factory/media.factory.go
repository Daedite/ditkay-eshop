package factory

import (
	"errors"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"strconv"
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
	var mediaId int
	if *body.Id != "" {
		id, err := strconv.Atoi(*body.Id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion %s", err))
		}
		mediaId = id
	}

	return &entity.Media{
		Id:          mediaId,
		Image:       *body.Image,
		Description: *body.Description,
	}, nil
}

var _ MediaFactory = &MediaFactoryImpl{}
