package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// LoadFilesToDB load to DB
func LoadFilesToDB(c echo.Context) error {
	return c.String(http.StatusOK, "loaded")
}
