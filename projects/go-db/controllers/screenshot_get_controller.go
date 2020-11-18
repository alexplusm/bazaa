package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ScreenshotGetController struct {
	ScreenshotCacheService interfaces.IScreenshotCacheService
	GameCacheService       interfaces.IGameCacheService
	UserService            interfaces.IUserService
}

func (controller *ScreenshotGetController) GetScreenshot(ctx echo.Context) error {
	gameID := ctx.Param("game-id")
	externalSystemID := ctx.QueryParam("extSystemId")
	userID := ctx.QueryParam("userId")

	if externalSystemID == "" {
		// BadRequest
		ctx.String(200, "kek")
		return nil
	}
	if userID == "" {
		// BadRequest
		ctx.String(200, "user required")
		return nil
	}

	ok := controller.GameCacheService.GameWithSameExtSystemIDExist(gameID, externalSystemID)
	if !ok {
		//return
	}
	// TODO: inject game service
	// TODO: BAD RESPONSE: game not started | game is finished | game not found

	fmt.Println("& GameWithSameExtSystemIDExist: ", ok)

	err := controller.UserService.CreateUser(userID)
	if err != nil {
		fmt.Println("USER SERVICE Error: ", err)
	}

	screenshot, hasScreenshot := controller.ScreenshotCacheService.GetScreenshot(gameID, userID)
	if !hasScreenshot {
		// TODO: game over
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	// TODO:!!!!!!
	res := struct {
		ScreenshotID string `json:"screenshot_id"`
		ImageURL     string `json:"image_url"`
	}{
		ScreenshotID: screenshot.ScreenshotID,
		ImageURL:     screenshot.ImageURL,
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(res))

	fmt.Printf("SCREEN: %+v\n", screenshot)
	fmt.Println("CONTEXT: ", gameID, externalSystemID, userID)

	return nil
}
