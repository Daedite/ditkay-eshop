package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeRepository interface {
	CreateMediaType(mediaType entity.MediaType) (models.MediaType, error)
	ReadMediaType(id int) (models.MediaType, error)
	UpdateMediaType(mediaType entity.MediaType) (models.MediaType, error)
	DeleteMediaType(mediaType entity.MediaType) (bool, error)
	ReadMediaTypes() ([]gorm.MediaType, error)
}
