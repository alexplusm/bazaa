package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type GameListController struct {
	GameService interfaces.IGameService
}

func (controller *GameListController) GetGames(ctx echo.Context) error {
	log.Info("in get games")
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)

	gamesBO, err := controller.GameService.GetGames()

	return nil
}
