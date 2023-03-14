package gorm

import (
	"fmt"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	gormModel "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"gorm.io/gorm"
)

type MediaRepositoryImpl struct {
	GormDB  *gorm.DB
	Factory factory.MediaFactory
}

func NewMediaRepositoryImpl(gormDB *gorm.DB, factory factory.MediaFactory) *MediaRepositoryImpl {
	return &MediaRepositoryImpl{
		GormDB:  gormDB,
		Factory: factory,
	}
}

var _ repository.MediaRepository = &MediaRepositoryImpl{}

func (m MediaRepositoryImpl) CreateMedia(media entity.Media) (models.Media, error) {
	var gormMedia *gormModel.Media = m.Factory.CreateMediaFactory(media).(*gormModel.Media)
	if err := m.GormDB.Create(&gormMedia).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create media id: %d, err: %s", media.Id, err))
		return nil, err
	}
	return gormMedia, nil
}

func (m MediaRepositoryImpl) ReadMedia(id string) (models.Media, error) {
	gormMedia := &gormModel.Media{}
	if err := m.GormDB.Where("id = ?", id).First(&gormMedia).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get media with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormMedia, nil
}

func (m MediaRepositoryImpl) UpdateMedia(media entity.Media) (models.Media, error) {
	var gormMedia *gormModel.Media = m.Factory.CreateMediaFactory(media).(*gormModel.Media)
	if err := m.GormDB.Where("id = ?", media.Id).Updates(media).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update media id: %d", media.Id))
		return nil, err
	}
	return gormMedia, nil
}

func (m MediaRepositoryImpl) DeleteMedia(media entity.Media) (bool, error) {
	gormMediaType := &gormModel.Media{}
	if err := m.GormDB.Where("id = ?", media.Id).Delete(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update media id: %d", media.Id))
		return false, err
	}
	return true, nil
}

func (m MediaRepositoryImpl) ReadMedias() ([]gormModel.Media, error) {
	gormMediaType := []gormModel.Media{}
	if err := m.GormDB.Find(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads media"))
		return nil, err
	}
	return gormMediaType, nil
}
