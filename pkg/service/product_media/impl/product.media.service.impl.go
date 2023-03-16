package impl

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product_media"
)

type ProductMediaServiceImpl struct {
	Repository   repository.ProductMediaRepository
	MediaService media.MediaService
}

func NewProductMediaService(repository repository.ProductMediaRepository, mediaService media.MediaService) *ProductMediaServiceImpl {
	return &ProductMediaServiceImpl{Repository: repository,
		MediaService: mediaService,
	}
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

func (p ProductMediaServiceImpl) ReadProductMediasByProductId(productId string) ([]entity.Media, error) {
	var medias []entity.Media
	mediaProducts, err := p.Repository.ReadProductMediasByProductId(productId)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	if len(mediaProducts) < 0 {
		return medias, nil
	}
	for _, mediaProduct := range mediaProducts {
		mediaResponse, err := p.MediaService.ReadMedia(mediaProduct.MediaId)
		if err != nil {
			continue
		}
		medias = append(medias, *mediaResponse)
	}
	return medias, err
}

func (p ProductMediaServiceImpl) RemoveMedias(mediaId string) (bool, error) {
	productMedia, err := p.Repository.ReadProductMediaWithImageId(mediaId)
	if err != nil {
		return false, err
	}
	pm := productMedia.GetProductMedia()
	if _, err := p.DeleteProductMedia(*pm); err != nil {
		return false, err
	}
	media, err := p.MediaService.ReadMedia(mediaId)
	if err != nil {
		return false, err
	}
	if _, err := p.MediaService.DeleteMedia(*media); err != nil {
		return false, err
	}

	return true, nil
}
