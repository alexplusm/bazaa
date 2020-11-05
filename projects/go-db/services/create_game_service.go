package services

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ICreateGameService interface {
	CreateGame()
}

type CreateGameService struct {
	Repository interfaces.IGameRepository
}

func (service *CreateGameService) CreateGame() {
	fmt.Println("CreateGame service")
	// pgx -> create game in DB
}
