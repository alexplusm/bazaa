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
	AnswerService     interfaces.IAnswerService
	GameService       interfaces.IGameService
	ScreenshotService interfaces.IScreenshotService
}

func (controller *ScreenshotResultsController) GetResult(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)
	screenshotID := ctx.Param(consts.ScreenshotIDUrlParam)

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

	screenshotExist, err := controller.ScreenshotService.ScreenshotExist(screenshotID)
	if err != nil {
		log.Error("get result: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !screenshotExist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("screenshot not found"),
		)
	}

	// TODO: In service
	res, err := controller.AnswerService.GetScreenshotResults(gameID, screenshotID)
	if err != nil {
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	usersAnswer := "undefined"
	for _, r := range res {
		if r.Result == "right" {
			usersAnswer = r.Answer
		}
	}

	resp := dto.ScreenshotResultsDTO{
		Finished:    len(res) == consts.RequiredAnswerCountToFinishScreenshot,
		Answers:     res,
		UsersAnswer: usersAnswer,
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
