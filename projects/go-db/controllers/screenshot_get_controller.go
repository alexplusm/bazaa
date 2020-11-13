package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ScreenshotGetController struct {
	ScreenshotCacheService interfaces.IScreenshotCacheService
	GameCacheService       interfaces.IGameCacheService
}

func (controller *ScreenshotGetController) GetScreenshot(ctx echo.Context) error {
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

	ok := controller.GameCacheService.GameWithSameExtSystemIDExist(gameID, externalSystemID)
	fmt.Println("OOOKKK: ", ok)

	// service.GameExist(gameID, externalSystemID) // TODO: check game exsitance and externalSystem
	// BAD RESPONSE -> game does not exist

	// TODO: BAD RESPONSE: game not started | game is finished | game not found

	// service.getScreenshot(gameID)

	fmt.Println("ctx", gameID, externalSystemID, userID)
	return nil
}
