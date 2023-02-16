package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeRepository interface {
	CreateMediaType(mediaType entity.MediaType) (models.MediaType, error)
	ReadMediaType(id string) (models.MediaType, error)
	UpdateMediaType(mediaType entity.MediaType) (models.MediaType, error)
	DeleteMediaType(mediaType entity.MediaType) (models.MediaType, error)
	ReadMediaTypes() ([]models.MediaType, error)
}
