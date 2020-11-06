package controllers

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UpdateGameController struct {
	Service interfaces.IUpdateGameService
}

func (service *UpdateGameController) UpdateGameController() {
	fmt.Println("Update controller")
}
