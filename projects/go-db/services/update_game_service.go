package services

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

type UpdateGameService struct {
	GameRepo       interfaces.IGameRepository
	SourceRepo     interfaces.ISourceRepository
	ScreenshotRepo interfaces.IScreenshotRepository
}

func (service *UpdateGameService) AttachZipArchiveToGame(gameID string, archives []*multipart.FileHeader) error {
	hasGame, err := service.GameRepo.HasNotStartedGameWithSameID(gameID)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	fmt.Println("has game: ", hasGame) // todo: if true -> return bad request

	filenames, err := fileutils.CopyFiles(archives, consts.MediaTempDir)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	images, err := fileutils.UnzipImages(filenames)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	source := dao.SourceDAO{dao.ArchiveSourceType, time.Now().Unix(), gameID}
	sourceID, err := service.SourceRepo.InsertSource(source)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	//fmt.Println("sourceID: ", sourceID)

	a, b := split(images, gameID, sourceID)

	err = service.ScreenshotRepo.InsertScreenshots(a)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	err = service.ScreenshotRepo.InsertScreenshotsWithExpertAnswer(b)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	//fmt.Println("images count", len(images))
	//fmt.Println("kek ||| ", len(a), len(b))
	//fmt.Printf("--- %+v\n", b[1])

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

	// INFO: Когда загружаем несколько архивов могут быть попасться одинаковые файлы
	// -> обрабатываем эту ситуацию
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
