package services

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type IGameService interface {
	Method1()
}

type GameService struct {
	interfaces.IGameRepository
}

func (service *GameService) Method1() {
	fmt.Println("Method1")
}
