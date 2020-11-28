package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type GameListController struct {
	GameService interfaces.IGameService
}

func (controller *GameListController) GetGames(ctx echo.Context) error {
	log.Info("in get games")

	controller.GameService.GetGames()

	return nil
}
