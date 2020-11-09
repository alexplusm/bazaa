package interfaces

import (
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type ICreateGameService interface {
	CreateGame(game bo.GameBO) (string, error)
}

type IUpdateGameService interface {
	AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error
	AttachSchedulesToGame(gameID string) error

	// TODO: move to -> GameService ?
	GetGame(gameID string) (bo.GameBO, error)
}
