package factory

import (
	"errors"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"strconv"
)

type ProductMediaFactory interface {
	CreateProductMedia(body spec.ProductMedia) (*entity.ProductMedia, error)
}
type ProductMediaFactoryImpl struct {
}

func NewProductMediaFactoryImpl() *ProductMediaFactoryImpl {
	return &ProductMediaFactoryImpl{}
}

var _ ProductMediaFactory = &ProductMediaFactoryImpl{}

func (m ProductMediaFactoryImpl) CreateProductMedia(body spec.ProductMedia) (*entity.ProductMedia, error) {
	var productMediaId int
	var mediaId int
	var mediaType int
	var productId int
	var err error
	if *body.Id != "" {
		id, err := strconv.Atoi(*body.Id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion Id: %s", err))
		}
		productMediaId = id
	}
	if *body.MediaId != "" {
		mediaId, err = strconv.Atoi(*body.MediaId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion MediaId: %s", err))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Media Id missing!"))
	}

	if *body.MediaType != "" {
		mediaType, err = strconv.Atoi(*body.MediaType)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion mediaType: %s", err))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("MediaType missing!"))
	}

	if *body.ProductId != "" {
		productId, err = strconv.Atoi(*body.ProductId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during conversion product id: %s", err))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("ProductId missing!"))
	}

	return &entity.ProductMedia{
		Id:        productMediaId,
		MediaId:   mediaId,
		MediaType: mediaType,
		ProductId: productId,
	}, nil
}
