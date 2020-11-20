package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ScreenshotSetAnswerController struct {
	ScreenshotCacheService      interfaces.IScreenshotCacheService
	ScreenshotUserAnswerService interfaces.IScreenshotUserAnswerService
}

func (controller *ScreenshotSetAnswerController) SetAnswer(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	screenshotID := ctx.Param("screenshot-id")

	userAnswerRaw := new(dto.UserAnswerRequestBody)
	if err := ctx.Bind(userAnswerRaw); err != nil {
		return fmt.Errorf("set answer controller: %v", err)
	}

	userAnswerBO := new(bo.UserAnswerBO)
	if err := userAnswerBO.Create(*userAnswerRaw, validate); err != nil {
		return fmt.Errorf("set answer controller: %v", err)
	}

	// TODO: check gameID and ExtSystemID

	screenshotExist := controller.ScreenshotCacheService.ScreenshotExist(screenshotID)
	if !screenshotExist {
		return ctx.String(200, "screenshot doesn't exist")
	}

	if !controller.ScreenshotCacheService.CanSetUserAnswerToScreenshot(
		userAnswerBO.UserID, screenshotID,
	) {
		// TODO: Что делать в этому случае? Обсудить с Колей
		// TODO: doc
		return ctx.String(200, "Can't Set UserAnswerToScreenshot")
	}

	answers, err := controller.ScreenshotCacheService.SetUserAnswer(
		userAnswerBO.UserID, screenshotID, userAnswerBO.Answer,
	)
	if err != nil {
	}

	screenshotIsFinished := controller.ScreenshotUserAnswerService.ScreenshotIsFinished(answers)
	if screenshotIsFinished {
		fmt.Println("UserID: ", userAnswerBO.UserID)
		controller.ScreenshotUserAnswerService.SaveUsersAnswers(answers, gameID, screenshotID)
		controller.ScreenshotCacheService.RemoveScreenshot(gameID, screenshotID)
	}

	response := controller.ScreenshotUserAnswerService.BuildUserAnswerResponse(userAnswerBO.UserID, answers)

	fmt.Printf("UserAnswer: %+v\n", *userAnswerBO)
	fmt.Println("Answers: ", response)
	fmt.Println("SetAnswer: ScreenshotID: ", screenshotID)
	fmt.Println()

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(response))
}
