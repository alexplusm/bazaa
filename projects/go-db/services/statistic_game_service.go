package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type StatisticGameService struct {
	ActiveUsersService interfaces.IActiveUsersService
	AnswerService      interfaces.IAnswerService
	ScreenshotService  interfaces.IScreenshotService
}

func (service *StatisticGameService) GetStatistics(games []bo.GameBO) (dto.StatisticGameDTO, error) {
	totalCount := 0
	answeredCount := 0
	activityUsers := 0
	usersList := make([]string, 0, 1024)

	for _, g := range games {
		c, _ := service.ScreenshotService.ScreenshotCountByGame(g.GameID)
		res, _ := service.AnswerService.GetUsersAndScreenshotCountByGame(g.GameID)
		answeredCount += res.Count
		totalCount += c
		usersList = append(usersList, res.UserID...)
		actUsers, _ := service.ActiveUsersService.CountOfActiveUsers(g.GameID)
		activityUsers += actUsers
	}

	usersMap := make(map[string]bool)
	for _, userID := range usersList {
		usersMap[userID] = true
	}

	resp := dto.StatisticGameDTO{
		ScreenshotsResolved: answeredCount,
		ScreenshotsLeft:     totalCount - answeredCount,
		UsersUnique:         len(usersMap),
		UsersActive:         activityUsers,
	}

	return resp, nil
}
