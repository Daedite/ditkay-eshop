package models

import "github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"

type Media interface {
	GetMedia() *entity.Media
}

type Product interface {
	GetProduct() *entity.Product
}

type ProductMedia interface {
	GetProductMedia() *entity.ProductMedia
}

type MediaType interface {
	GetMediaType() *entity.MediaType
}
