package product

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product_media"
)

type ProductServiceImpl struct {
	Repository          repository.ProductRepository
	ProductMediaService product_media.ProductMediaService
}

func NewProductServiceImpl(
	repository repository.ProductRepository,
	productMediaService product_media.ProductMediaService,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		Repository:          repository,
		ProductMediaService: productMediaService,
	}
}

var _ product.ProductService = ProductServiceImpl{}

func (p ProductServiceImpl) CreateProduct(product entity.Product) (*entity.Product, error) {
	prod, err := p.Repository.CreateProduct(product)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProduct(), err
}

func (p ProductServiceImpl) ReadProduct(id string) (*entity.Product, error) {
	prod, err := p.Repository.ReadProduct(id)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProduct(), err
}

func (p ProductServiceImpl) UpdateProduct(product entity.Product) (*entity.Product, error) {
	prod, err := p.Repository.UpdateProduct(product)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return prod.GetProduct(), err
}

func (p ProductServiceImpl) DeleteProduct(product entity.Product) (bool, error) {
	prod, err := p.Repository.DeleteProduct(product)
	if err != nil {
		logger.Log.Error(err.Error())
		return false, err
	}
	return prod, err
}

func (p ProductServiceImpl) ReadProducts() (*[]entity.Product, error) {
	prod, err := p.Repository.ReadProducts()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return getMediaList(prod), err
}

func getMediaList(list []gorm.Product) *[]entity.Product {
	var products []entity.Product
	if len(list) < 0 {
		return &products
	}
	for _, Product := range list {
		products = append(products, entity.Product{
			Id:          Product.Id,
			Name:        Product.Name,
			Description: Product.Description,
			Quantity:    Product.Quantity,
			SellPrice:   Product.SellPrice,
			BuyPrice:    Product.BuyPrice,
		})
	}
	return &products
}

func (p ProductServiceImpl) ReadProductsForClient() (*[]entity.ProductClient, error) {
	products, err := p.ReadProducts()
	if err != nil {
		return nil, err
	}
	var productClient []entity.ProductClient
	for _, product := range *products {
		//TODO we will need to get only the selected productMedia.
		medias, err := p.ProductMediaService.ReadProductMediasByProductId(product.Id)
		if err != nil {
			continue
		}
		productClient = append(productClient, entity.ProductClient{Media: medias[0], Product: product})
	}
	return &productClient, nil
}
