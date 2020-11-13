package interfaces

import (
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type IGameService interface {
	CreateGame(game bo.GameBO) (string, error)
	GetGame(gameID string) (bo.GameBO, error)
}

type IAttachSourceToGameService interface {
	AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error
	AttachSchedulesToGame(gameID string) error
}

type IExtSystemService interface {
	CreateExtSystem(extSystem bo.ExtSystemBO) error
}

type IGameCacheService interface {
	PrepareGame(gameID string)
	GameWithSameExtSystemIDExist(gameID, extSystemID string) bool
}

type IScreenshotCacheService interface {
}
