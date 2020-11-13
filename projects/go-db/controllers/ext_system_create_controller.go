package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

type ExtSystemCreateController struct {
}

func (controller *ExtSystemCreateController) CreateExtSystem(ctx echo.Context) error {
	fmt.Println("CreateExtSystem controller")
	return nil
}
