package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/models"
)

type IGameRepository interface {
	CreateGame(game models.GameModel) (string, error)
}
