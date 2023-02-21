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

func (d DitKayEshopApiController) DeleteMedia(ctx echo.Context) error {
	var media spec.Media
	logger.Log.Info("MediaId received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.MediaFactory.CreateMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaService.DeleteMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (d DitKayEshopApiController) PatchMedia(ctx echo.Context) error {
	var media spec.Media
	logger.Log.Info("MediaId received for update.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.MediaFactory.CreateMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaService.UpdateMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (d DitKayEshopApiController) PostMedia(ctx echo.Context) error {
	var media spec.Media
	logger.Log.Info("MediaId received for creation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	mediaEntity, err := d.MediaFactory.CreateMedia(media)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	response, err := d.MediaService.CreateMedia(*mediaEntity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (d DitKayEshopApiController) GetMediaMediaId(ctx echo.Context, mediaId string) error {

	if mediaId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(mediaId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("error during conversion %s", err))
	}
	response, err := d.MediaService.ReadMedia(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (d DitKayEshopApiController) GetMedias(ctx echo.Context) error {
	response, err := d.MediaService.ReadMedias()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, response)
}
