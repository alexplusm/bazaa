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

type ScreenshotController struct {
	ScreenshotCacheService interfaces.IScreenshotCacheService
	GameCacheService       interfaces.IGameCacheService
	UserService            interfaces.IUserService
	ImageService           interfaces.IImageService
}

func (controller *ScreenshotController) Retrieve(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	qp := ScreenshotRetrieveQP{}
	qp.fromCtx(ctx)

	if qp.ExtSystemID.Value == "" {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem required"),
		)
	}
	if qp.UserID.Value == "" {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("user required"),
		)
	}

	ok := controller.GameCacheService.GameWithSameExtSystemIDExist(gameID, qp.ExtSystemID.Value)
	if !ok {
		//return
	}
	// TODO: inject game service
	// TODO: BAD RESPONSE: game not started | game is finished | game not found

	err := controller.UserService.CreateUser(qp.UserID.Value)
	if err != nil {
		log.Error("Error log")
		fmt.Println("USER SERVICE Error: ", err)
	}

	screenshot, hasScreenshot := controller.ScreenshotCacheService.GetScreenshot(gameID, qp.UserID.Value)
	if !hasScreenshot {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game is finished"),
		)
	}

	// TODO: screenshot.ImageURL -> ImageName
	// TODO: into service ScreenshotCacheService.Retrieve
	imageURL, err := controller.ImageService.BuildImageURL(screenshot.ImageURL)
	if err != nil {
		log.Error("get screenshot controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	res := struct {
		ScreenshotID string `json:"screenshot_id"`
		ImageURL     string `json:"image_url"`
	}{
		ScreenshotID: screenshot.ScreenshotID,
		ImageURL:     imageURL,
	}

	fmt.Println("Get screenshot: ", " | ", screenshot.ScreenshotID)

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(res))
}
