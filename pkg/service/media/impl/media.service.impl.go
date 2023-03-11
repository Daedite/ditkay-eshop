package impl

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media"
)

type MediaServiceImpl struct {
	Repository repository.MediaRepository
}

func NewMediaService(repository repository.MediaRepository) *MediaServiceImpl {
	return &MediaServiceImpl{Repository: repository}
}

var _ media.MediaService = MediaServiceImpl{}

func (m MediaServiceImpl) CreateMedia(media entity.Media) (*entity.Media, error) {
	med, err := m.Repository.CreateMedia(media)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return med.GetMedia(), nil
}

func (m MediaServiceImpl) ReadMedia(id string) (*entity.Media, error) {
	med, err := m.Repository.ReadMedia(id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return med.GetMedia(), nil
}

func (m MediaServiceImpl) UpdateMedia(media entity.Media) (*entity.Media, error) {
	med, err := m.Repository.UpdateMedia(media)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return med.GetMedia(), nil
}

func (m MediaServiceImpl) DeleteMedia(media entity.Media) (bool, error) {
	med, err := m.Repository.DeleteMedia(media)
	if err != nil {
		logger.Log.Error(err.Error())
		return false, err
	}
	return med, nil
}

func (m MediaServiceImpl) ReadMedias() (*[]entity.Media, error) {
	med, err := m.Repository.ReadMedias()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return getMediaList(med), nil
}

func getMediaList(list []gorm.Media) *[]entity.Media {
	var medias []entity.Media
	if len(list) < 0 {
		return &medias
	}
	for _, media := range list {
		medias = append(medias, entity.Media{
			Id:          media.Id,
			Image:       *media.Image,
			Description: *media.Description,
		})
	}
	return &medias
}
