package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type ExtSystemCreateController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *ExtSystemCreateController) CreateExtSystem(ctx echo.Context) error {
	extSystemRaw := new(dto.CreateExtSystemRequestBody)

	if err := ctx.Bind(extSystemRaw); err != nil {
		return fmt.Errorf("extSystem create controller: %v\n", err)
	}

	extSystem := new(bo.ExtSystemBO)
	if err := extSystem.FromDTO(*extSystemRaw, validate); err != nil {
		return fmt.Errorf("extSystem create controller: %v\n", err)
	}

	fmt.Printf("ExtSystemRaw: %+v\n", *extSystemRaw)
	fmt.Printf("ExtSystem: %+v | hasID: %v\n", extSystem, extSystem.HasID())

	return controller.ExtSystemService.CreateExtSystem(*extSystem)
}
