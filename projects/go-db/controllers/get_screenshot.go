package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

type GetScreenshotController struct {
}

func (controller *GetScreenshotController) GetScreenshot(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	externalSystemID := ctx.QueryParam("extSystemId")
	userID := ctx.QueryParam("userId")

	if externalSystemID == "" {
		ctx.String(200, "kek")
		return nil
	}
	if userID == "" {
		ctx.String(200, "user required")
		return nil
	}

	// TODO: check game exsitance and externalSystem

	// service.GameExist(gameID, externalSystemID)
	// BAD RESPONSE -> game does not exist

	// TODO: BAD RESPONSE: game not started | game is finished | game not found

	// TODO: all params - required

	fmt.Println("ctx", gameID, externalSystemID, userID)
	return nil
}
