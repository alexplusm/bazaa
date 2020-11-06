package interfaces

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type ICreateGameService interface {
	CreateGame(game bo.GameBO) (string, error)
}

type IUpdateGameService interface {
	AttachZipArchiveToGame()
	AttachSchedulesToGame()
}
