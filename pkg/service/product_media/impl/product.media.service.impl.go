package impl

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product_media"
)

type ProductMediaServiceImpl struct {
	Repository repository.ProductMediaRepository
}

func NewProductMediaService(repository repository.ProductMediaRepository) *ProductMediaServiceImpl {
	return &ProductMediaServiceImpl{Repository: repository}
}

var _ product_media.ProductMediaService = &ProductMediaServiceImpl{}

func (p ProductMediaServiceImpl) CreateProductMedia(productMedia entity.ProductMedia) (*entity.ProductMedia, error) {
	prod, err := p.Repository.CreateProductMedia(productMedia)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProductMedia(), err
}

func (p ProductMediaServiceImpl) ReadProductMedia(id string) (*entity.ProductMedia, error) {
	prod, err := p.Repository.ReadProductMedia(id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProductMedia(), err
}

func (p ProductMediaServiceImpl) UpdateProductMedia(productMedia entity.ProductMedia) (*entity.ProductMedia, error) {
	prod, err := p.Repository.UpdateProductMedia(productMedia)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProductMedia(), err
}

func (p ProductMediaServiceImpl) DeleteProductMedia(productMedia entity.ProductMedia) (bool, error) {
	prod, err := p.Repository.DeleteProductMedia(productMedia)
	if err != nil {
		logger.Log.Error(err.Error())
		return false, err
	}
	return prod, err
}

func (p ProductMediaServiceImpl) ReadProductMedias() (*[]entity.ProductMedia, error) {
	prod, err := p.Repository.ReadProductMedias()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return getProductMediaList(prod), err
}

func getProductMediaList(list []gorm.ProductMedia) *[]entity.ProductMedia {
	var products []entity.ProductMedia
	if len(list) < 0 {
		return &products
	}
	for _, Product := range list {
		products = append(products, entity.ProductMedia{
			Id:        Product.Id,
			MediaType: Product.MediaId,
			ProductId: Product.ProductId,
			MediaId:   Product.MediaId,
		})
	}
	return &products
}

func (p ProductMediaServiceImpl) ReadProductMediasByProductId(productId string) (*[]entity.ProductMedia, error) {
	prod, err := p.Repository.ReadProductMediasByProductId(productId)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return getProductMediaList(prod), err
}
