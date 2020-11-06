package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type IGameRepository interface {
	CreateGame(game dao.GameDAO) (string, error)
}

type ISourceRepository interface {
	CreateSource(source dao.SourceDAO) error
}
