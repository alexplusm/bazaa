package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type IGameService interface {
	CreateGame(game bo.GameBO) (string, error)
}
