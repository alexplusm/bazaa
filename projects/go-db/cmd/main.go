package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Alexplusm/bazaa/projects/go-db/configs"
	"github.com/Alexplusm/bazaa/projects/go-db/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/infrastructures"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/files"
)

/* source: https://github.com/irahardianto/service-pattern-go */

func main() {
	defer infrastructures.Injector().CloseStoragesConnections()

	initDirs()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	registerRoutes(e)

	// TODO: PORT from .env
	// TODO: use own logger?
	e.Logger.Fatal(e.Start(":1234"))
}

func initDirs() {
	dirs := []string{configs.MediaRoot, configs.MediaTempDir}
	for _, dir := range dirs {
		files.CreateDirIfNotExists(dir)
	}
}

func registerRoutes(e *echo.Echo) {
	container := infrastructures.Injector()

	// TODO: rename?
	createGameController := container.InjectCreateGameController()

	// if err := middlewares.ContentTypeMiddleware(ctx, "application/json"); err != nil {
	// 	return err
	// }
	e.POST("api/v1/game", createGameController.CreateGame)

	// g := e.Group("api/v1/game")
	// g.Use(middle2)

	e.GET("/check/alive", controllers.ItsAlive)

	// e.PUT("api/v1/game/:game-id", controllers.UpdateGame(conn))

	// TODO: wait refactoring
	//e.POST("api/v1/game", func(ctx echo.Context) error {
	//	// TODO: Groups and middlewares
	//	return controllers.CreateGame(conn)(ctx)
	//})

}

// ------ test
//func middle1(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		if err := next(ctx); err != nil {
//			ctx.Error(err)
//		}
//		fmt.Println("middle 1")
//		return nil
//	}
//}
//
//func middle2(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		if err := next(ctx); err != nil {
//			ctx.Error(err)
//		}
//		fmt.Println("middle 2")
//		return nil
//	}
//}
