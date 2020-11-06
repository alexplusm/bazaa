package services

import (
	"fmt"
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
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

	filenames, err := fileutils.CopyFiles(archives, consts.MediaTempDir)
	if err != nil {
		return fmt.Errorf("attach zip archive: %+v\n", err)
	}

	// must return imageNames and category
	res, err := fileutils.UnzipImages(filenames)
	if err != nil {
		return fmt.Errorf("attach zip archive: %+v\n", err)
	}

	fmt.Println("filenames:", filenames)
	fmt.Println("Has game:", hasGame)
	fmt.Println("FILES", res, "| len:", len(res))
	fmt.Println()

	// todo: fill database use res

	removeArchives(filenames)

	return nil
}

func (service *UpdateGameService) AttachSchedulesToGame(gameID string) error {
	// TODO:later
	fmt.Println("Schedules attaching coming soon ... : gameID =", gameID)
	return nil
}

func removeArchives(filenames []string) {
	for _, fn := range filenames {
		if err := fileutils.RemoveFile(consts.MediaTempDir, fn); err != nil {
			fmt.Println(err) // TODO:log // TODO:error
		}
	}
}
