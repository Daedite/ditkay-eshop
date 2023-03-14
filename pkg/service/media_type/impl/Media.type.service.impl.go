package impl

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media_type"
)

type MediaTypeServiceImpl struct {
	Repository repository.MediaTypeRepository
}

func NewMediaTypeServiceImpl(repository repository.MediaTypeRepository) *MediaTypeServiceImpl {
	return &MediaTypeServiceImpl{Repository: repository}
}

var _ media_type.MediaTypeService = MediaTypeServiceImpl{}

func (m MediaTypeServiceImpl) CreateMediaType(mediaType entity.MediaType) (*entity.MediaType, error) {
	prod, err := m.Repository.CreateMediaType(mediaType)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetMediaType(), err
}

func (m MediaTypeServiceImpl) ReadMediaType(id int) (*entity.MediaType, error) {
	prod, err := m.Repository.ReadMediaType(id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetMediaType(), err
}

func (m MediaTypeServiceImpl) UpdateMediaType(mediaType entity.MediaType) (*entity.MediaType, error) {
	prod, err := m.Repository.UpdateMediaType(mediaType)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetMediaType(), err
}

func (m MediaTypeServiceImpl) DeleteMediaType(mediaType entity.MediaType) (bool, error) {
	prod, err := m.Repository.DeleteMediaType(mediaType)
	if err != nil {
		logger.Log.Error(err.Error())
		return false, err
	}
	return prod, err
}

func (m MediaTypeServiceImpl) ReadMediaTypes() (*[]entity.MediaType, error) {
	medias, err := m.Repository.ReadMediaTypes()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return getMediaTypeList(medias), err
}

func getMediaTypeList(list []gorm.MediaType) *[]entity.MediaType {
	var mediaTypes []entity.MediaType
	if len(list) < 0 {
		return &mediaTypes
	}
	for _, mediaType := range list {
		mediaTypes = append(mediaTypes, entity.MediaType{
			Id:          mediaType.Id,
			Name:        mediaType.Name,
			Description: mediaType.Description,
		})
	}
	return &mediaTypes
}
