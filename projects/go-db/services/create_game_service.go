package services

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/models"
)

type ICreateGameService interface {
	CreateGame()
}

type CreateGameService struct {
	Repository interfaces.IGameRepository
}

func (service *CreateGameService) CreateGame(game models.GameModel) (string, error) {
	fmt.Printf("CreateGame service: %+v\n", game)

	return service.Repository.CreateGame(game)
}
