package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/domain"
)

type IGameService interface {
	CreateGame(game domain.GameBO) (string, error)
}
