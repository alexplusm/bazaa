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

	fmt.Println("has game: ", hasGame) // todo: if true -> return bad request

	filenames, err := fileutils.CopyFiles(archives, consts.MediaTempDir)
	if err != nil {
		return fmt.Errorf("attach zip archive: %+v\n", err)
	}

	images, err := fileutils.UnzipImages(filenames)
	if err != nil {
		return fmt.Errorf("attach zip archive: %+v\n", err)
	}

	// --- TODO: Когда загружаем несколько архивов
	// TODO: могут быть несколько одинаковых файлов -> обработать!
	//
	//mmap := make(map[string]int)
	//for _, r := range res {
	//	mmap[r.Filename]++
	//}
	//for key := range mmap {
	//	if mmap[key] > 1 {
	//		fmt.Println("RESULT: ", mmap[key])
	//	}
	//}
	// ---

	//images -> db
	//source -> db

	fmt.Println("images count", len(images))

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
