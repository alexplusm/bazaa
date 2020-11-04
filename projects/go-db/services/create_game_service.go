package services

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type IGameService interface {
	CreateGame()
}

type GameService struct {
	interfaces.IGameRepository
}

func (service *GameService) CreateGame() {
	fmt.Println("CreateGame service")
	// pgx -> create game in DB
}
