package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsGameController struct {
	GameService        interfaces.IGameService
	ExtSystemService   interfaces.IExtSystemService
	ScreenshotService  interfaces.IScreenshotService
	AnswerService      interfaces.IAnswerService
	ActiveUsersService interfaces.IActiveUsersService
}

func (controller *StatisticsGameController) GetStatistics(ctx echo.Context) error {
	qp := dto.StatisticsUserQueryParams{}
	qp.FromCTX(ctx)

	exist, err := controller.ExtSystemService.ExtSystemExist(qp.ExtSystemID)
	if err != nil {
		log.Error("get games controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("extSystem not found"),
		)
	}

	games, err := controller.GameService.GetGames(qp.ExtSystemID)

	expectedGames := make([]bo.GameBO, 0, len(games))

	// filter games
	if len(qp.GameIDs) != 0 {
		for _, game := range games {
			for _, gameQP := range qp.GameIDs {
				if game.GameID == gameQP {
					expectedGames = append(expectedGames, game)
				}
			}
		}
	} else {
		expectedGames = games
	}

	totalCount := 0
	answeredCount := 0
	activityUsers := 0
	usersList := make([]string, 0, 1024)

	for _, g := range expectedGames {
		c, _ := controller.ScreenshotService.ScreenshotCountByGame(g.GameID)
		res, _ := controller.AnswerService.GetUsersAndScreenshotCountByGame(g.GameID)
		answeredCount += res.Count
		totalCount += c
		usersList = append(usersList, res.UserID...)
		actUsers, _ := controller.ActiveUsersService.CountOfActiveUsers(g.GameID)
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

	return ctx.JSON(
		http.StatusOK,
		httputils.BuildSuccessResponse(resp),
	)
}
