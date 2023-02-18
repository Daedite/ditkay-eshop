package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

// ProductMedia defines model for ProductMedia.
type ProductMedia struct {
	Id        int `gorm:"primarykey"`
	MediaId   int `sql:"media_id"`
	MediaType int `sql:"media_type"`
	ProductId int `sql:"product_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p ProductMedia) GetProductMedia() *entity.ProductMedia {
	return &entity.ProductMedia{
		Id:        p.Id,
		MediaId:   p.MediaId,
		MediaType: p.MediaType,
		ProductId: p.ProductId,
	}
}

var _ models.ProductMedia = ProductMedia{}

func (ProductMedia) TableName() string {
	return "product_media"
}
