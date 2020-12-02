package controllers

import (
	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ExtSystemListController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *ExtSystemListController) List(ctx echo.Context) error {
	return nil
}
