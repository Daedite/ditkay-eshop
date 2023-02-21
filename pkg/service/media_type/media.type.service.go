package media_type

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
)

type MediaTypeService interface {
	CreateMediaType(mediaType entity.MediaType) (*entity.MediaType, error)
	ReadMediaType(id int) (*entity.MediaType, error)
	UpdateMediaType(mediaType entity.MediaType) (*entity.MediaType, error)
	DeleteMediaType(mediaType entity.MediaType) (bool, error)
	ReadMediaTypes() (*[]entity.MediaType, error)
}
