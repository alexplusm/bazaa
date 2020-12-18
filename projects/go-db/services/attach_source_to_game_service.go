package services

import (
	"archive/zip"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/fileutils"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/logutils"
)

type AttachSourceToGameService struct {
	GameRepo           interfaces.IGameRepo
	ScreenshotRepo     interfaces.IScreenshotRepo
	SourceService      interfaces.ISourceService
	FileService        interfaces.IFileService
	ImageFilterService interfaces.IImageFilterService
}

func (service *AttachSourceToGameService) AttachArchives(
	gameID string, archives []*multipart.FileHeader,
) error {
	archivesPaths, err := service.FileService.SaveFiles(archives, consts.MediaTempDir)
	if err != nil {
		return fmt.Errorf("%v AttachArchives: %v", logutils.GetStructName(service), err)
	}

	files, err := service.FileService.UnzipArchives(archivesPaths, consts.MediaRoot)
	if err != nil {
		return fmt.Errorf("%v AttachArchives: %v", logutils.GetStructName(service), err)
	}

	files = service.ImageFilterService.Filter(files)

	value := getStr(archives)

	sourceID, err := service.SourceService.Create(gameID, value, consts.ArchiveSourceType)
	if err != nil {
		return fmt.Errorf("%v AttachArchives: %v", logutils.GetStructName(service), err)
	}

	// TODO: refactor
	newImg := parseImageCategory(files)

	a, b := split(newImg, gameID, sourceID)

	err = service.ScreenshotRepo.InsertList(a)
	if err != nil {
		return fmt.Errorf("%v AttachArchives: %v", logutils.GetStructName(service), err)
	}

	err = service.ScreenshotRepo.InsertListWithExpertAnswer(b)
	if err != nil {
		return fmt.Errorf("%v AttachArchives: %v", logutils.GetStructName(service), err)
	}

	fileutils.RemoveFiles(archivesPaths)

	return nil
}

func (service *AttachSourceToGameService) AttachSchedules(gameID string) error {
	fmt.Println("Schedules attaching coming soon ... : gameID =", gameID)
	return nil
}

func (service *AttachSourceToGameService) AttachGameResults(gameID string, params bo.AttachGameParams) error {
	sourceID, err := service.SourceService.Create(gameID, params.SourceGameID, consts.AnotherGameResultSourceType)
	if err != nil {
		return fmt.Errorf("%v AttachGameResults: %v", logutils.GetStructName(service), err)
	}

	screenshots, err := service.ScreenshotRepo.SelectListByGameID(params.SourceGameID)
	if err != nil {
		return fmt.Errorf("%v AttachGameResults: %v", logutils.GetStructName(service), err)
	}

	newScreenshots := make([]dao.ScreenshotDAO, 0, len(screenshots))

	for _, screenshot := range screenshots {
		if string(screenshot.UsersAnswer) == params.Answer {
			ddao := dao.ScreenshotDAO{
				Filename: screenshot.Filename,
				GameID:   gameID,
				SourceID: sourceID,
			}
			newScreenshots = append(newScreenshots, ddao)
		}
	}

	err = service.ScreenshotRepo.InsertList(newScreenshots)
	if err != nil {
		return fmt.Errorf("%v AttachGameResults: %v", logutils.GetStructName(service), err)
	}

	return nil
}

// todo: не интересно - желательно удалить
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
				screen := dao.ScreenshotDAO{Filename: image.Filename, GameID: gameID, SourceID: sourceID}

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

// TODO: rename | replace in another package ?
func getStr(archives []*multipart.FileHeader) string {
	archivesFilename := make([]string, 0, len(archives))

	for _, archive := range archives {
		archivesFilename = append(archivesFilename, archive.Filename)
	}

	return strings.Join(archivesFilename, ",")
}

func parseImageCategory(files []zip.File) []bo.ImageParsingResult {
	results := make([]bo.ImageParsingResult, 0, len(files))

	for _, file := range files {
		fname := file.FileInfo().Name()
		withViolation := strings.HasSuffix(file.Name, filepath.Join(withViolationDir, fname))
		noViolation := strings.HasSuffix(file.Name, filepath.Join(noViolationDir, fname))

		var result bo.ImageParsingResult

		if withViolation {
			result = bo.ImageParsingResult{Filename: fname, Category: WithViolationCategory}
		} else if noViolation {
			result = bo.ImageParsingResult{Filename: fname, Category: NoViolationCategory}
		} else {
			result = bo.ImageParsingResult{Filename: fname, Category: UndefinedCategory}
		}

		results = append(results, result)
	}

	return results
}
