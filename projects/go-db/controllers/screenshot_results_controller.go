package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	//log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ScreenshotResultsController struct {
	AnswerService interfaces.IAnswerService
}

func (controller *ScreenshotResultsController) GetResult(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	screenshotID := ctx.Param("screenshot-id")

	fmt.Println(gameID, screenshotID)

	//dto.ScreenshotResultsDTO{}
	//[]dto.UserAnswerForScreenshotResultDTO{}
	controller.AnswerService.GetScreenshotResults(gameID, screenshotID)

	// TODO: screenshot exist | game exist

	return nil
}
