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

	screenshotExist := controller.ScreenshotCacheService.ScreenshotExist(screenshotID)
	if !screenshotExist {
		ctx.String(200, "screenshot doesn't exist")
		return nil
	}

	if !controller.ScreenshotCacheService.CanSetUserAnswerToScreenshot(
		userAnswerBO.UserID, screenshotID,
	) {
		// TODO: Что делать в этому случае?
		// TODO: Обсудить с Колей
		ctx.String(200, "Can't Set UserAnswerToScreenshot")
		return nil
	}

	controller.ScreenshotCacheService.SetUserAnswerToScreenshot(
		userAnswerBO.UserID, screenshotID, userAnswerBO.Answer,
	)
	answers := controller.ScreenshotCacheService.GetUsersAnswers(screenshotID)

	// if len(answers) == required... -> culc response | write in db!
	// TODO: check count of answers

	ctx.JSON(
		http.StatusOK,
		// TODO: getData -> in service
		httputils.BuildResponse(getData(userAnswerBO.UserID, answers)),
	)

	fmt.Printf("UserAnswer: %+v\n", *userAnswerBO)
	fmt.Println("Answers: ", answers)
	fmt.Println("SetAnswer: Params: ", gameID, screenshotID)

	return nil
}

func getData(userID string, answersBO []bo.UserAnswerCacheBO) dto.UserAnswerResponseData {
	answers := make([]dto.UserAnswerKek, 0, len(answersBO))
	finished := false

	// TODO: вычислять результат (в сервисе!)
	for _, answer := range answersBO {
		answerDTO := dto.UserAnswerKek{
			UserID: answer.UserID, Answer: answer.Answer, Result: "rees",
		}
		answers = append(answers, answerDTO)
	}

	return dto.UserAnswerResponseData{Finished: finished, UserResult: "us-res", Answers: answers}
}
