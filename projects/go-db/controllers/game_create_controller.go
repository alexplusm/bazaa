package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameCreateController struct {
	GameService interfaces.IGameService
}

func (controller *GameCreateController) CreateGame(ctx echo.Context) error {
	gameRaw := new(dto.CreateGameRequestBody)

	if err := ctx.Bind(gameRaw); err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK, httputils.BuildBadRequestErrorResponse(),
		)
	}

	game := new(bo.GameBO)
	if err := game.CreateGame(*gameRaw, validate); err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("validation"),
		)
	}

	gameID, err := controller.GameService.CreateGame(*game)
	if err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK, httputils.BuildBadRequestErrorResponse(),
		)
	}

	response := httputils.BuildSuccessResponse(dto.CreateGameResponseBody{GameID: gameID})

	return ctx.JSON(http.StatusOK, response)
}
