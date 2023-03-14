package controller

import (
	"encoding/json"
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (d DitKayEshopApiController) DeleteMediaType(ctx echo.Context) error {
	var media spec.MediaType
	logger.Log.Info("MediaType received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaTypeEntity, err := d.MediaTypeFactory.CreateMediaType(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaTypeService.DeleteMediaType(*mediaTypeEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) PatchMediaType(ctx echo.Context) error {
	var media spec.MediaType
	logger.Log.Info("MediaType received for updating.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaTypeEntity, err := d.MediaTypeFactory.CreateMediaType(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaTypeService.UpdateMediaType(*mediaTypeEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) PostMediaType(ctx echo.Context) error {
	var media spec.MediaType
	logger.Log.Info("MediaType received for creation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaTypeEntity, err := d.MediaTypeFactory.CreateMediaType(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaTypeService.CreateMediaType(*mediaTypeEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetMediaTypeMediaTypeId(ctx echo.Context, mediaTypeId string) error {
	if mediaTypeId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(mediaTypeId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("error during conversion %s", err))
	}
	response, err := d.MediaTypeService.ReadMediaType(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (d DitKayEshopApiController) GetMediaTypes(ctx echo.Context) error {
	response, err := d.MediaTypeService.ReadMediaTypes()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
