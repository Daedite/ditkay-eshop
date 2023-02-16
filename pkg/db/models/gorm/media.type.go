package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

type MediaType struct {
	Description *string `gorm:"primarykey"`
	Id          *string
	Name        *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m MediaType) GetMediaType() *entity.MediaType {
	return &entity.MediaType{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
	}
}

var _ models.MediaType = MediaType{}

func (MediaType) TableName() string {
	return "media_type"
}
