package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"gorm.io/gorm"
)

type MediaRepositoryImpl struct {
	GormDB *gorm.DB
}

func NewMediaRepositoryImpl(gormDB *gorm.DB) *MediaRepositoryImpl {
	return &MediaRepositoryImpl{
		GormDB: gormDB,
	}
}

var _ repository.MediaRepository = &MediaRepositoryImpl{}

func (m MediaRepositoryImpl) CreateMedia(media entity.Media) (models.Media, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaRepositoryImpl) ReadMedia(id string) (models.Media, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaRepositoryImpl) UpdateMedia(media entity.Media) (models.Media, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaRepositoryImpl) DeleteMedia(media entity.Media) (models.Media, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaRepositoryImpl) ReadMedias() ([]models.Media, error) {
	//TODO implement me
	panic("implement me")
}
