package services

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UpdateGameService struct {
	GameRepo   interfaces.IGameRepository
	SourceRepo interfaces.ISourceRepository
}

func (service *UpdateGameService) AttachZipArchiveToGame(gameID string) error {
	hasGame, err := service.GameRepo.HasHotStartedGameWithSameID(gameID)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	fmt.Println("Has game:", hasGame)

	return nil
}

func (service *UpdateGameService) AttachSchedulesToGame() error {
	// TODO:later
	return nil
}
