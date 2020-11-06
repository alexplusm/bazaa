package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type IGameRepository interface {
	InsertGame(game dao.GameDAO) (string, error)
	HasNotStartedGameWithSameID(gameID string) (bool, error)
}

type ISourceRepository interface {
	CreateSource(source dao.SourceDAO) error
}
