package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UpdateGameService struct {
	GameRepo   interfaces.IGameRepository
	SourceRepo interfaces.ISourceRepository
}

func (service *UpdateGameService) AttachZipArchiveToGame() {}

func (service *UpdateGameService) AttachSchedulesToGame() {}
