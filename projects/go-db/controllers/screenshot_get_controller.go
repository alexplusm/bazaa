package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ScreenshotGetController struct {
	ScreenshotCacheService interfaces.IScreenshotCacheService
	GameCacheService       interfaces.IGameCacheService
	UserService            interfaces.IUserService
	ImageService           interfaces.IImageService
}

func (controller *ScreenshotGetController) GetScreenshot(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	// TODO: queryParams
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

	err := controller.UserService.CreateUser(userID)
	if err != nil {
		log.Error("Error log")
		fmt.Println("USER SERVICE Error: ", err)
	}

	screenshot, hasScreenshot := controller.ScreenshotCacheService.GetScreenshot(gameID, userID)
	if !hasScreenshot {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game is finished"),
		)
	}

	// TODO:!!!!!!

	// TODO: screenshot.ImageURL -> ImageName
	imageURL := controller.ImageService.BuildImageURL(screenshot.ImageURL)

	res := struct {
		ScreenshotID string `json:"screenshot_id"`
		ImageURL     string `json:"image_url"`
	}{
		ScreenshotID: screenshot.ScreenshotID,
		ImageURL:     imageURL,
	}

	fmt.Println("Get screenshot: ", userID, " | ", screenshot.ScreenshotID)

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(res))
}
