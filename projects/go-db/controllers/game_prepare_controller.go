package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type GamePrepareController struct {
	GameCacheService interfaces.IGameCacheService
}

func (controller *GamePrepareController) PrepareGame(ctx echo.Context) error {
	log.Debug("PrepareGame", ctx)
	return nil
}
