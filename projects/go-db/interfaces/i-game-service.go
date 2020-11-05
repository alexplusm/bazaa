package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/models"
)

type IGameService interface {
	CreateGame(game models.GameModel) (string, error)
}
