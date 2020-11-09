package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// ItsAlive server alive check
func ItsAlive(c echo.Context) error {
	msg := "*** I'm ALIVE! ***"
	fmt.Println(msg)
	return c.String(http.StatusOK, msg)
}
