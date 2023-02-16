package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeFactory struct {
}

func (m MediaTypeFactory) CreateMediaTypeFactory(mediaType entity.MediaType) models.MediaType {
	return gorm.MediaType{
		Id:          mediaType.Id,
		Name:        mediaType.Name,
		Description: mediaType.Description,
	}
}

func NewMediaTypeFactory() *MediaTypeFactory {
	return &MediaTypeFactory{}
}

var _ factory.MediaTypeFactory = &MediaTypeFactory{}
