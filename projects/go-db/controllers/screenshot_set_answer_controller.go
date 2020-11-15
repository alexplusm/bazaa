package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type ScreenshotSetAnswerController struct {
	ScreenshotCacheService interfaces.IScreenshotCacheService
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

	controller.ScreenshotCacheService.SetUserAnswerToScreenshot(
		userAnswerBO.UserID, screenshotID, userAnswerBO.Answer,
	)

	fmt.Printf("UserAnswer: %+v\n", *userAnswerBO)
	fmt.Println("SetAnswer: Params: ", gameID, screenshotID)

	return nil
}
