package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaFactory struct {
}

func (m MediaFactory) CreateMediaFactory(media entity.Media) models.Media {
	return &gorm.Media{
		Id:          media.Id,
		Description: media.Description,
		Image:       media.Image,
	}
}

var _ factory.MediaFactory = &MediaFactory{}
