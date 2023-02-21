package media

import "github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"

type MediaService interface {
	CreateMedia(media entity.Media) (*entity.Media, error)
	ReadMedia(id int) (*entity.Media, error)
	UpdateMedia(media entity.Media) (*entity.Media, error)
	DeleteMedia(media entity.Media) (bool, error)
	ReadMedias() (*[]entity.Media, error)
}
