package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ScreenshotSetAnswerController struct {
	RedisClient interfaces.IRedisHandler
}

func (controller *ScreenshotSetAnswerController) SetAnswer(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	screenshotID := ctx.Param("screenshot-id")

	fmt.Println("SetAnswer: Params: ", gameID, screenshotID)
	return nil
}
