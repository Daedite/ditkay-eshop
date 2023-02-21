package factory

import (
	"errors"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"strconv"
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
	var mediaTypeId int
	if *body.Id != "" {
		id, err := strconv.Atoi(*body.Id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion %s", err))
		}
		mediaTypeId = id
	}

	return &entity.MediaType{
		Id:          mediaTypeId,
		Name:        *body.Name,
		Description: *body.Description,
	}, nil
}

var _ MediaTypeFactory = &MediaTypeFactoryImpl{}
