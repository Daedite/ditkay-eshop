package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (d DitKayEshopApiController) DeleteProductMedia(ctx echo.Context) error {
	var media spec.ProductMedia
	logger.Log.Info("Product Media received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.ProductMediaFactory.CreateProductMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductMediaService.DeleteProductMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) PatchProductMedia(ctx echo.Context) error {
	var media spec.ProductMedia
	logger.Log.Info("Product Media received for update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.ProductMediaFactory.CreateProductMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductMediaService.UpdateProductMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetProductMedias(ctx echo.Context) error {
	logger.Log.Info("Product Media request received for Read all operation.")

	response, err := d.ProductMediaService.ReadProductMedias()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) PostProductMedia(ctx echo.Context) error {
	var media spec.ProductMedia
	logger.Log.Info("Product Media received for create operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.ProductMediaFactory.CreateProductMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.ProductMediaService.CreateProductMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetProductMediaProductMediaId(ctx echo.Context, productMediaId string) error {
	var ProductMedia int
	logger.Log.Info("Product Media received for Read operation.")
	if productMediaId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("Missing required param."))
	}
	ProductMedia, err := strconv.Atoi(productMediaId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.New(fmt.Sprintf("error converting product media id %s", err)))
	}

	response, err := d.ProductMediaService.ReadProductMedia(ProductMedia)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
