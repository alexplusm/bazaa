package main

import (
	"github.com/Alexplusm/bazaa/projects/go-db/src/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/src/dbcon"
	"github.com/labstack/echo"
)

func main() {
	dbcon.Connect()

	e := echo.New()

	e.POST("/upload/images", controllers.UploadFiles)

	// todo: PORT from .env
	e.Logger.Fatal(e.Start(":1234"))
}
