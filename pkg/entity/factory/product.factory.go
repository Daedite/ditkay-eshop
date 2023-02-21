package factory

import (
	"errors"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"strconv"
)

type ProductFactory interface {
	CreateProduct(body spec.Product) (*entity.Product, error)
}
type ProductFactoryImpl struct {
}

func NewProductFactoryImpl() *ProductFactoryImpl {
	return &ProductFactoryImpl{}
}

func (m ProductFactoryImpl) CreateProduct(body spec.Product) (*entity.Product, error) {
	var productId int
	if *body.Id != "" {
		id, err := strconv.Atoi(*body.Id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion %s", err))
		}
		productId = id
	}
	if *body.Name == "" {
		return nil, errors.New(fmt.Sprintf("Product name required."))
	}
	if *body.SellPrice < 0 {
		return nil, errors.New(fmt.Sprintf("Product price required."))
	}

	return &entity.Product{
		Id:          productId,
		Name:        *body.Name,
		BuyPrice:    *body.BuyPrice,
		SellPrice:   *body.SellPrice,
		Quantity:    *body.Quantity,
		Description: *body.Description,
	}, nil
}

var _ ProductFactory = &ProductFactoryImpl{}
