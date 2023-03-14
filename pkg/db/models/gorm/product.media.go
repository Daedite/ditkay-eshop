package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

// ProductMedia defines model for ProductMedia.
type ProductMedia struct {
	Id        string `gorm:"primarykey"`
	MediaId   string `sql:"media_id"`
	MediaType string `sql:"media_type"`
	ProductId string `sql:"product_id"`
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
