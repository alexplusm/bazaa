package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/infrastructures"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

/*
*	source:
*		https://github.com/irahardianto/service-pattern-go
*		https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508
 */

const (
	errorPrefix = "main: "
)

func main() {
	setupLogger()

	injector, err := infrastructures.Injector()
	if err != nil {
		log.Fatal(errorPrefix, err)
	}
	defer injector.CloseStoragesConnections()

	initDirs()

	e := echo.New()

	e.Use(middleware.Logger())

	e.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		// TODO: func !
		usernameAdm := []byte(os.Getenv("SERVER_ADMIN_USERNAME"))
		passwordAdm := []byte(os.Getenv("SERVER_ADMIN_PASSWORD"))

		if subtle.ConstantTimeCompare([]byte(username), usernameAdm) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), passwordAdm) == 1 {
			return true, nil
		}

		return false, ctx.JSON(http.StatusUnauthorized, nil) // TODO: body
	}))

	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		log.Error(errorPrefix, err)
	}

	err = registerRoutes(e)
	if err != nil {
		log.Fatal(errorPrefix, err)
	}

	log.Fatal(e.Start(":" + os.Getenv("SERVER_PORT_INNER")))
}

func initDirs() {
	dirs := []string{consts.MediaRoot, consts.MediaTempDir}
	for _, dir := range dirs {
		fileutils.CreateDirIfNotExists(dir)
	}
}

// TODO
//func registerMiddlewares() {}

func registerRoutes(e *echo.Echo) error {
	injector, err := infrastructures.Injector()
	if err != nil {
		return fmt.Errorf("register routes: %v", err)
	}

	gameController := injector.InjectGameController()

	gamePrepareController := injector.InjectGamePrepareController()

	extSystemController := injector.InjectExtSystemController()

	screenshotController := injector.InjectScreenshotController()
	screenshotSetAnswerController := injector.InjectScreenshotSetAnswerController()
	screenshotResultsController := injector.InjectScreenshotResultsController()

	statisticsUserController := injector.InjectStatisticUserController()
	statisticsLeaderboardController := injector.InjectStatisticLeaderboardController()
	statisticsGameController := injector.InjectStatisticGameController()

	// TODO:later
	// Create middleware for each route with whitelist of ContentTypes:
	// ["multipart/form-data"] | ["application/json"]

	// TODO: ["application/json"]
	e.POST("api/v1/game", gameController.Create)

	e.GET("api/v1/game", gameController.List)

	e.GET("api/v1/game/:"+consts.GameIDUrlParam, gameController.Details)

	// TODO: ["multipart/form-data"]
	e.PUT("api/v1/game/:"+consts.GameIDUrlParam+"/archives", gameController.AttachArchives)

	// TODO: ["application/json"]
	e.PUT("api/v1/game/:"+consts.GameIDUrlParam+"/schedules", gameController.AttachSchedules)

	// TODO: ["application/json"]
	e.PUT("api/v1/game/:"+consts.GameIDUrlParam+"/game-results", gameController.AttachGameResults)

	// TODO: ["application/json"]
	e.POST("api/v1/game/prepare", gamePrepareController.PrepareGame)

	// TODO: ["application/json"]
	e.POST("api/v1/ext_system", extSystemController.Create)

	e.GET("api/v1/ext_system", extSystemController.List)

	e.GET("api/v1/game/:"+consts.GameIDUrlParam+"/screenshot", screenshotController.Retrieve)

	// TODO: ["application/json"]
	e.POST(
		"api/v1/game/:"+consts.GameIDUrlParam+"/screenshot/:"+consts.ScreenshotIDUrlParam+"/answer",
		screenshotSetAnswerController.SetAnswer,
	)
	e.GET("/api/v1/game/:"+consts.GameIDUrlParam+"/screenshot/:"+consts.ScreenshotIDUrlParam+"/result", screenshotResultsController.GetResult)

	e.GET("api/v1/statistics/user/:"+consts.UserIDUrlParam, statisticsUserController.GetStatistics)

	e.GET("api/v1/statistics/users/leaderboard", statisticsLeaderboardController.GetStatistics)

	e.GET("api/v1/statistics/games", statisticsGameController.GetStatistics)

	e.GET("api/v1/check/alive", controllers.ItsAlive)

	return nil
}

func setupLogger() {
	// source: https://www.honeybadger.io/blog/golang-logging/
	// TODO: log in file if PROD
	// log.SetOutput()
	log.SetFormatter(&log.JSONFormatter{})
}
