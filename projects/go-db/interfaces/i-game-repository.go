package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/dao"
)

type IGameRepository interface {
	CreateGame(game dao.GameDAO) (string, error)
}
