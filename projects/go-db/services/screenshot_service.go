package services

import "github.com/Alexplusm/bazaa/projects/go-db/interfaces"

type ScreenshotService struct {
	ScreenshotRepo interfaces.IScreenshotRepo
}

func (repo *ScreenshotService) Exist(screenshotID string) (bool, error) {
	return repo.ScreenshotRepo.Exist(screenshotID)
}

func (repo *ScreenshotService) ScreenshotCountByGame(gameID string) (int, error) {
	return repo.ScreenshotRepo.CountByGame(gameID)
}
