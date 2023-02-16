package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/factory"
	gormModel "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
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
	var gormMediaType *gormModel.MediaType = m.Factory.CreateMediaTypeFactory(mediaType).(*gormModel.MediaType)
	if err := m.GormDB.Create(gormMediaType).Error; err != nil {
		return nil, err
	}
	return gormMediaType, nil
}

func (m MediaTypeRepositoryImpl) ReadMediaType(id string) (models.MediaType, error) {
	gormMediaType := &gormModel.MediaType{}
	if err := m.GormDB.First(&gormMediaType, id).Error; err != nil {
		return nil, err
	}
	return gormMediaType, nil
}

func (m MediaTypeRepositoryImpl) UpdateMediaType(mediaType entity.MediaType) (models.MediaType, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaTypeRepositoryImpl) DeleteMediaType(mediaType entity.MediaType) (models.MediaType, error) {
	//TODO implement me
	panic("implement me")
}

func (m MediaTypeRepositoryImpl) ReadMediaTypes() ([]models.MediaType, error) {
	//TODO implement me
	panic("implement me")
}
