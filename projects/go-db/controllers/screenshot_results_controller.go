package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ScreenshotResultsController struct {
	AnswerService interfaces.IAnswerService
	GameService   interfaces.IGameService
}

func (controller *ScreenshotResultsController) GetResult(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	screenshotID := ctx.Param("screenshot-id")

	fmt.Println(gameID, screenshotID)

	gameExist, err := controller.GameService.GameExist(gameID)
	if err != nil {
		log.Error("get result: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !gameExist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}

	// TODO: screenshot exist | game exist

	res, err := controller.AnswerService.GetScreenshotResults(gameID, screenshotID)
	if err != nil {
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	resp := dto.ScreenshotResultsDTO{
		Finished: len(res) == consts.RequiredAnswerCountToFinishScreenshot,
		Answers:  res,
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
