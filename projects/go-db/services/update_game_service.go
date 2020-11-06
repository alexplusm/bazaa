package services

import (
	"fmt"
	"mime/multipart"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

type UpdateGameService struct {
	GameRepo   interfaces.IGameRepository
	SourceRepo interfaces.ISourceRepository
}

func (service *UpdateGameService) AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error {
	hasGame, err := service.GameRepo.HasNotStartedGameWithSameID(gameID)
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

	// source -> db : sourceID
	// sourceID, gameID -> images -> db

	a, b := split(images, gameID, "sourceID")

	fmt.Println("images count", len(images))
	fmt.Println("kek ||| ", len(a), len(b))
	fmt.Printf("--- %+v\n", b[1])

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

func split(images []fileutils.ImageParsingResult, gameID, sourceID string) ([]dao.ScreenshotDAO, []dao.ScreenshotWithExpertAnswerDAO) {
	mmap := make(map[string]bool)
	imagesWithoutExpertAnswer := make([]dao.ScreenshotDAO, 0, len(images))
	imagesWithExpertAnswer := make([]dao.ScreenshotWithExpertAnswerDAO, 0, len(images))

	for _, image := range images {
		if !mmap[image.Filename] {
			mmap[image.Filename] = true
			if image.Category == fileutils.UndefinedCategory {
				screen := dao.ScreenshotDAO{image.Filename, gameID, sourceID}

				imagesWithoutExpertAnswer = append(imagesWithoutExpertAnswer, screen)
			} else {
				screen := dao.ScreenshotWithExpertAnswerDAO{
					dao.ScreenshotDAO{image.Filename, gameID, sourceID},
					image.Category,
				}
				imagesWithExpertAnswer = append(imagesWithExpertAnswer, screen)
			}
		}
	}

	return imagesWithoutExpertAnswer, imagesWithExpertAnswer
}
