package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/infrastructures"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

/* source: https://github.com/irahardianto/service-pattern-go */

func main() {
	defer infrastructures.Injector().CloseStoragesConnections()

	initDirs()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		fmt.Printf("Error: %v\n", err)
	}

	registerRoutes(e)

	// TODO: PORT from .env
	// TODO: use own logger?
	e.Logger.Fatal(e.Start(":1234"))
}

func initDirs() {
	dirs := []string{consts.MediaRoot, consts.MediaTempDir}
	for _, dir := range dirs {
		fileutils.CreateDirIfNotExists(dir)
	}
}

func registerRoutes(e *echo.Echo) {
	injector := infrastructures.Injector()

	gameCreateController := injector.InjectGameCreateController()
	gameUpdateController := injector.InjectGameUpdateController()
	extSystemCreateController := injector.InjectExtSystemCreateController()
	screenshotGetController := injector.InjectScreenshotGetController()

	// TODO:later
	// Create middleware for each route with whitelist of ContentTypes:
	// ["application/json", "multipart/form-data"] | ["application/json"]

	// TODO: ["application/json"]
	e.POST("api/v1/game", gameCreateController.CreateGame)
	// TODO: ["application/json", "multipart/form-data"]
	e.PUT("api/v1/game/:game-id", gameUpdateController.UpdateGame)
	// TODO: ["application/json"]
	e.POST("api/v1/ext-system", extSystemCreateController.CreateExtSystem)
	// TODO: ["application/json"]
	e.GET("api/v1/game/:game-id/screenshot", screenshotGetController.GetScreenshot)

	// TODO: for test
	e.GET("/check/alive", controllers.ItsAlive)

	testService(injector)
}

func testService(i infrastructures.IInjector) {
	gameCacheService := i.InjectGameCacheService()
	gameID := "baabf15b-3a05-4592-9935-101637c12d67"
	gameCacheService.PrepareGame(gameID)
}
