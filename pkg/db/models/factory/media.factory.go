package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaFactory interface {
	CreateMediaFactory(media entity.Media) models.Media
}
