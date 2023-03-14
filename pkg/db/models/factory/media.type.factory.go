package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeFactory interface {
	CreateMediaTypeFactory(mediaType entity.MediaType) models.MediaType
}
