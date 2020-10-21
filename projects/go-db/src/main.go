package main

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/src/controllers"
	"github.com/Alexplusm/bazaa/projects/go-db/src/dbcon"
	"github.com/Alexplusm/bazaa/projects/go-db/src/models"
	"github.com/labstack/echo"

	"github.com/jackc/pgx/pgxpool"
)

func main() {
	conn, err := dbcon.Connect()
	if err != nil {
		// logger
		fmt.Println("Unable to connection to database: %v\n", err)
	}
	defer conn.Close()
	fmt.Println("Connected to database!")

	dbTest(conn) // test

	e := echo.New()

	e.POST("/upload/images", controllers.UploadFiles)
	e.GET("/upload/check", controllers.LoadFilesToDBWrapper(conn))

	// todo: PORT from .env
	e.Logger.Fatal(e.Start(":1234"))
}

func dbTest(p *pgxpool.Pool) {
	img := models.ImageDao{URL: "url/kekus", Category: "no"}
	models.InsertImage(p, img)
}
