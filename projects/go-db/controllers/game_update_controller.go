package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameUpdateController struct {
	GameService               interfaces.IGameService
	AttachSourceToGameService interfaces.IAttachSourceToGameService
}

func (controller *GameUpdateController) UpdateGame(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	switch httputils.ParseContentType(ctx) {
	case consts.FormDataContentType:
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
		}

		game, err := controller.GameService.GetGame(gameID)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildErrorResponse(http.StatusOK, "game not found"),
			)
		}

		if !game.NotStarted() {
			log.Info("game update controller: ", "game started: ", gameID)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildErrorResponse(http.StatusOK, "game started"),
			)
		}

		archives := form.File["archives"]

		err = controller.AttachSourceToGameService.AttachZipArchiveToGame(gameID, archives)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildBadRequestErrorResponse(),
			)
		}

		// TODO: return gameID?
		return ctx.JSON(http.StatusOK, httputils.BuildSuccessWithoutBodyResponse())
	case consts.ApplicationContentJSON:
		err := controller.AttachSourceToGameService.AttachSchedulesToGame(gameID)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildBadRequestErrorResponse(),
			)
		}
	}

	return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
}
