package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GamePrepareController struct {
	GameCacheService interfaces.IGameCacheService
}

func (controller *GamePrepareController) PrepareGame(ctx echo.Context) error {
	prepareGame := new(dto.PrepareGameResponseBody)

	if err := ctx.Bind(prepareGame); err != nil {
		log.Error("prepare game controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	err := controller.GameCacheService.PrepareGame(prepareGame.GameID)
	if err != nil {
		log.Error("prepare game controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	log.Info("prepare game success: ", prepareGame.GameID)
	return ctx.JSON(http.StatusOK, httputils.BuildSuccessWithoutBodyResponse())
}
