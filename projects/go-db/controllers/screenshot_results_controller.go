package controllers

import (
	"fmt"

	"github.com/labstack/echo"
	//log "github.com/sirupsen/logrus"
)

type ScreenshotResultsController struct {
}

func (controller *ScreenshotResultsController) GetResult(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	screenshotID := ctx.Param("screenshot-id")

	fmt.Println(gameID, screenshotID)

	// TODO: screenshot exist | game exist

	return nil
}
