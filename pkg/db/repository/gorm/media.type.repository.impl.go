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

type MediaTypeRepositoryImpl struct {
	GormDB  *gorm.DB
	Factory factory.MediaTypeFactory
}

func NewMediaTypeRepositoryImpl(gormDB *gorm.DB, factory factory.MediaTypeFactory) *MediaTypeRepositoryImpl {
	return &MediaTypeRepositoryImpl{
		GormDB:  gormDB,
		Factory: factory,
	}
}

var _ repository.MediaTypeRepository = &MediaTypeRepositoryImpl{}

func (m MediaTypeRepositoryImpl) CreateMediaType(mediaType entity.MediaType) (models.MediaType, error) {
	var gormMediaType gormModel.MediaType = m.Factory.CreateMediaTypeFactory(mediaType).(gormModel.MediaType)
	if err := m.GormDB.Create(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to create Media type Name: %s", mediaType.Name))
		return nil, err
	}
	return gormMediaType, nil
}

func (m MediaTypeRepositoryImpl) ReadMediaType(id int) (models.MediaType, error) {
	gormMediaType := &gormModel.MediaType{}
	if err := m.GormDB.First(&gormMediaType, id).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to get Media type id: %d", id))
		return nil, err
	}
	return gormMediaType, nil
}

func (m MediaTypeRepositoryImpl) UpdateMediaType(mediaType entity.MediaType) (models.MediaType, error) {
	var gormMediaType gormModel.MediaType = m.Factory.CreateMediaTypeFactory(mediaType).(gormModel.MediaType)
	if err := m.GormDB.Model(&gormModel.MediaType{}).Where("id = ? ", mediaType.Id).Updates(gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update mediaType id: %d", mediaType.Id))
		return nil, err
	}
	return gormMediaType, nil
}

func (m MediaTypeRepositoryImpl) DeleteMediaType(mediaType entity.MediaType) (bool, error) {
	gormMediaType := &gormModel.MediaType{}
	if err := m.GormDB.Where("id = ?", mediaType.Id).Delete(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to delete mediaType id: %d", mediaType.Id))
		return false, err
	}
	return true, nil
}

func (m MediaTypeRepositoryImpl) ReadMediaTypes() ([]gormModel.MediaType, error) {
	gormMediaType := []gormModel.MediaType{}
	if err := m.GormDB.Find(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads mediaType"))
		return nil, err
	}
	return gormMediaType, nil
}
