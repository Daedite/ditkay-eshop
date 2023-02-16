package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

type Media struct {
	Id          *string `gorm:"primarykey"`
	Image       *string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m Media) GetMedia() *entity.Media {
	return &entity.Media{
		Id:          m.Id,
		Image:       m.Image,
		Description: m.Description,
	}
}

var _ models.Media = Media{}

func (Media) TableName() string {
	return "media"
}
