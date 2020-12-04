package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ItsAlive(c echo.Context) error {
	msg := "*** I'm ALIVE! ***"
	fmt.Println(msg)
	return c.String(http.StatusOK, msg)
}
