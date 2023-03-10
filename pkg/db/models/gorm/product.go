package gorm

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"time"
)

type Product struct {
	BuyPrice    float32 `gorm:"primarykey"`
	Description string
	Id          string
	Name        string
	Quantity    int
	SellPrice   float32 `sql:"sell_price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p Product) GetProduct() *entity.Product {
	return &entity.Product{
		Id:          p.Id,
		Name:        p.Name,
		Quantity:    p.Quantity,
		SellPrice:   p.SellPrice,
		BuyPrice:    p.BuyPrice,
		Description: p.Description,
	}
}

var _ models.Product = Product{}

func (Product) TableName() string {
	return "product"
}
