package services

import "github.com/Alexplusm/bazaa/projects/go-db/interfaces"

type ScreenshotService struct {
	ScreenshotRepo interfaces.IScreenshotRepository
}

func (repo *ScreenshotService) ScreenshotExist(screenshotID string) (bool, error) {
	return repo.ScreenshotRepo.ScreenshotExist(screenshotID)
}
