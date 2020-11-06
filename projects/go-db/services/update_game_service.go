package services

import (
	"fmt"
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UpdateGameService struct {
	GameRepo   interfaces.IGameRepository
	SourceRepo interfaces.ISourceRepository
}

func (service *UpdateGameService) AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error {
	hasGame, err := service.GameRepo.HasHotStartedGameWithSameID(gameID)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	for _, archive := range archives {
		fmt.Println(archive.Filename)
	}

	fmt.Println("Has game:", hasGame)

	return nil
}

func (service *UpdateGameService) AttachSchedulesToGame(gameID string) error {
	// TODO:later
	fmt.Println("Schedules attaching coming soon ... : gameID =", gameID)
	return nil
}

//func removeArchives(filenames []string) {
//	for _, fn := range filenames {
//		if err := fileutils.RemoveFile(consts.MediaTempDir, fn); err != nil {
//			fmt.Println(err) // TODO:log // TODO:error
//		}
//	}
//}
