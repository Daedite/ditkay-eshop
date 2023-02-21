package product

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
}

func NewProductServiceImpl(repository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{Repository: repository}
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

func (p ProductServiceImpl) ReadProduct(id int) (*entity.Product, error) {
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
