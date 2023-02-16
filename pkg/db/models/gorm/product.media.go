package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

// ProductMedia defines model for ProductMedia.
type ProductMedia struct {
	Id          *string `gorm:"primarykey"`
	MediaId     *string `sql:"media_id"`
	MediaTypeId *string `sql:"media_type_id"`
	ProductId   *string `sql:"product_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p ProductMedia) GetProductMedia() *entity.ProductMedia {
	return &entity.ProductMedia{
		Id:        p.Id,
		MediaId:   p.MediaId,
		MediaType: p.MediaTypeId,
		ProductId: p.ProductId,
	}
}

var _ models.ProductMedia = ProductMedia{}

func (ProductMedia) TableName() string {
	return "product_media"
}
