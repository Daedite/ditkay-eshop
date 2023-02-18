package repository

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaRepository interface {
	CreateMedia(media entity.Media) (models.Media, error)
	ReadMedia(id int) (models.Media, error)
	UpdateMedia(media entity.Media) (models.Media, error)
	DeleteMedia(media entity.Media) (bool, error)
	ReadMedias() ([]gorm.Media, error)
}
