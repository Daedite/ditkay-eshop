package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductControllerImpl struct {
}

func (d DitKayEshopApiController) PatchProduct(ctx echo.Context) error {
	var product spec.Product
	logger.Log.Info("Product received for Updating.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createProduct, err := d.ProductFactory.CreateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductService.UpdateProduct(*createProduct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) PostProduct(ctx echo.Context) error {
	var product spec.Product
	logger.Log.Info("Product received Create request.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createProduct, err := d.ProductFactory.CreateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductService.CreateProduct(*createProduct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetProductProductId(ctx echo.Context, productId string) error {
	if productId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("productId missing error"))
	}
	response, err := d.ProductService.ReadProduct(productId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetProducts(ctx echo.Context) error {
	response, err := d.ProductService.ReadProducts()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) DeleteProduct(ctx echo.Context) error {
	var product spec.Product
	logger.Log.Info("Product received for Deleting.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createProduct, err := d.ProductFactory.CreateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductService.DeleteProduct(*createProduct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
