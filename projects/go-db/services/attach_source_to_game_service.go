package services

import (
	"fmt"
	"mime/multipart"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
)

type AttachSourceToGameService struct {
	GameRepo       interfaces.IGameRepository
	SourceRepo     interfaces.ISourceRepository
	ScreenshotRepo interfaces.IScreenshotRepository
	FileService    interfaces.IFileService
}

func (service *AttachSourceToGameService) AttachZipArchiveToGame(
	gameID string, archives []*multipart.FileHeader,
) error {
	filenames, err := service.FileService.CopyFiles(archives, consts.MediaTempDir)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	images, err := service.FileService.UnzipImages(filenames)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	source := dao.SourceDAO{Type: dao.ArchiveSourceType, CreatedAt: time.Now().Unix(), GameID: gameID}
	sourceID, err := service.SourceRepo.InsertOne(source)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	a, b := split(images, gameID, sourceID)
	err = service.ScreenshotRepo.InsertList(a)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	err = service.ScreenshotRepo.InsertListWithExpertAnswer(b)
	if err != nil {
		return fmt.Errorf("attach zip archive: %v", err)
	}

	removeArchives(filenames)

	return nil
}

func (service *AttachSourceToGameService) AttachSchedulesToGame(gameID string) error {
	fmt.Println("Schedules attaching coming soon ... : gameID =", gameID)
	return nil
}

func removeArchives(filenames []string) {
	for _, fn := range filenames {
		err := fileutils.RemoveFile(consts.MediaTempDir, fn)
		if err != nil {
			log.Error("remove archive: ", err)
		}
	}
}

func split(
	images []bo.ImageParsingResult, gameID, sourceID string,
) ([]dao.ScreenshotDAO, []dao.ScreenshotWithExpertAnswerDAO) {
	mmap := make(map[string]bool)
	imagesWithoutExpertAnswer := make([]dao.ScreenshotDAO, 0, len(images))
	imagesWithExpertAnswer := make([]dao.ScreenshotWithExpertAnswerDAO, 0, len(images))

	// INFO: Когда загружаем несколько архивов могут быть попасться одинаковые файлы
	// -> обрабатываем эту ситуацию
	for _, image := range images {
		if !mmap[image.Filename] {
			mmap[image.Filename] = true
			if image.Category == UndefinedCategory {
				screen := dao.ScreenshotDAO{image.Filename, gameID, sourceID}

				imagesWithoutExpertAnswer = append(imagesWithoutExpertAnswer, screen)
			} else {
				screen := dao.ScreenshotWithExpertAnswerDAO{
					ScreenshotDAO: dao.ScreenshotDAO{
						Filename: image.Filename, GameID: gameID, SourceID: sourceID,
					},
					ExpertAnswer: image.Category,
				}
				imagesWithExpertAnswer = append(imagesWithExpertAnswer, screen)
			}
		}
	}

	return imagesWithoutExpertAnswer, imagesWithExpertAnswer
}
