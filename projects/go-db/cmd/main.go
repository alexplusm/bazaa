package main

import (
	"fmt"
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
	e.Pre(middleware.RemoveTrailingSlash())
	// TODO: https://echo.labstack.com/middleware/logger
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

func registerRoutes(e *echo.Echo) error {
	injector, err := infrastructures.Injector()
	if err != nil {
		return fmt.Errorf("register routes: %v", err)
	}

	gameCreateController := injector.InjectGameCreateController()
	gameUpdateController := injector.InjectGameUpdateController()
	gamePrepareController := injector.InjectGamePrepareController()
	gameListController := injector.InjectGameListController()
	extSystemCreateController := injector.InjectExtSystemCreateController()

	screenshotGetController := injector.InjectScreenshotGetController()
	screenshotSetAnswerController := injector.InjectScreenshotSetAnswerController()
	screenshotResultsController := injector.InjectScreenshotResultsController()

	statisticsUserController := injector.InjectStatisticsUserController()
	statisticsLeaderboardController := injector.InjectStatisticsLeaderboardController()
	statisticsGameController := injector.InjectStatisticsGameController()

	// TODO:later
	// Create middleware for each route with whitelist of ContentTypes:
	// ["application/json", "multipart/form-data"] | ["application/json"]

	// TODO: ["application/json"]
	e.POST("api/v1/game", gameCreateController.CreateGame)

	e.GET("api/v1/game", gameListController.GetGames)

	// TODO: ["application/json"]
	e.POST("api/v1/game/prepare", gamePrepareController.PrepareGame)

	// TODO: ["application/json", "multipart/form-data"]
	e.PUT("api/v1/game/:game-id", gameUpdateController.UpdateGame)

	// TODO: ["application/json"]
	e.POST("api/v1/ext_system", extSystemCreateController.CreateExtSystem)

	e.GET("api/v1/game/:game-id/screenshot", screenshotGetController.GetScreenshot)

	// TODO: ["application/json"]
	e.POST(
		"api/v1/game/:game-id/screenshot/:screenshot-id/answer",
		screenshotSetAnswerController.SetAnswer,
	)
	e.GET("/api/v1/game/:game-id/screenshot/:screenshot-id/result", screenshotResultsController.GetResult)

	e.GET("api/v1/statistics/user/:user-id", statisticsUserController.GetStatistics)
	e.GET("api/v1/statistics/users/leaderboard", statisticsLeaderboardController.GetStatistics)
	e.GET("api/v1/statistics/games", statisticsGameController.GetStatistics)

	// TODO: for test
	e.GET("check/alive", controllers.ItsAlive)

	return nil
}

func setupLogger() {
	// source: https://www.honeybadger.io/blog/golang-logging/
	// TODO: log in file if PROD
	// log.SetOutput()
	log.SetFormatter(&log.JSONFormatter{})
}
