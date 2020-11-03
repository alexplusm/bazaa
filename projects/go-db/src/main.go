package main

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/src/configs"
	"github.com/Alexplusm/bazaa/projects/go-db/src/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/src/dbcon"
	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/files"
)

func main() {
	initDirs()

	conn, err := dbcon.Connect()
	if err != nil {
		// logger
		fmt.Printf("Unable to connection to database: %v\n", err)
	}
	defer conn.Close()
	fmt.Println("Connected to database!")

	// dbcon.RedisConnect()

	// todo: REMOVE TRAILING SLASH IN URLS (Rewrite midddleware in "e.Pre()")

	e := echo.New()

	// g := e.Group("api/v1/game")
	// g.Use(middle2)
	// e.Use(middle1)
	// e.Use(middle2)

	e.GET("/upload/images/test", controllers.LoadFilesToDBWrapper(conn))
	e.GET("/check/alive", controllers.ItsAlive)

	e.POST("api/v1/game", func(ctx echo.Context) error {
		// // TODO: Groups and middlewares
		// if err := middlewares.ContentTypeMiddleware(ctx, "application/json"); err != nil {
		// 	return err
		// }
		return controllers.CreateGame(conn)(ctx)
	})
	e.PUT("api/v1/game/:game-id", controllers.UpdateGame(conn))

	// todo: PORT from .env
	// todo: use own logger ?
	e.Logger.Fatal(e.Start(":1234"))
}

func initDirs() {
	dirs := []string{configs.MediaRoot, configs.MediaTempDir}
	for _, dir := range dirs {
		files.CreateDirIfNotExists(dir)
	}
}

// ------ test

func middle1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := next(ctx); err != nil {
			ctx.Error(err)
		}
		fmt.Println("middle 1")
		return nil
	}
}

func middle2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := next(ctx); err != nil {
			ctx.Error(err)
		}
		fmt.Println("middle 2")
		return nil
	}
}
