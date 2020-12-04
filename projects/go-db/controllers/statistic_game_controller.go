package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticGameController struct {
	ExtSystemService     interfaces.IExtSystemService
	GameService          interfaces.IGameService
	StatisticGameService interfaces.IStatisticGameService

	// TODO: delete after test
	//ScreenshotService  interfaces.IScreenshotService
	//AnswerService      interfaces.IAnswerService
	//ActiveUsersService interfaces.IActiveUsersService
}

func (controller *StatisticGameController) GetStatistics(ctx echo.Context) error {
	qp := StatisticGameQP{}
	qp.fromCtx(ctx)

	if qp.ExtSystemID.Value != "" {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem required"),
		)
	}

	exist, err := controller.ExtSystemService.ExtSystemExist(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("statistic game controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("extSystem not found"),
		)
	}

	games, err := controller.GameService.GetGames(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("statistic game controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	games = controller.GameService.FilterGames(qp.GameIDs.Value, games)
	if len(games) == 0 {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}

	// TODO: error?
	resp, _ := controller.StatisticGameService.GetStatistics(games)

	return ctx.JSON(
		http.StatusOK,
		httputils.BuildSuccessResponse(resp),
	)

	// TODO: service.GetStatistic()...
	// AFTER TEST DELETE
	//totalCount := 0
	//answeredCount := 0
	//activityUsers := 0
	//usersList := make([]string, 0, 1024)
	//
	//for _, g := range games {
	//	c, _ := controller.ScreenshotService.ScreenshotCountByGame(g.GameID)
	//	res, _ := controller.AnswerService.GetUsersAndScreenshotCountByGame(g.GameID)
	//	answeredCount += res.Count
	//	totalCount += c
	//	usersList = append(usersList, res.UserID...)
	//	actUsers, _ := controller.ActiveUsersService.CountOfActiveUsers(g.GameID)
	//	activityUsers += actUsers
	//}
	//
	//usersMap := make(map[string]bool)
	//for _, userID := range usersList {
	//	usersMap[userID] = true
	//}
	//
	//resp := dto.StatisticGameDTO{
	//	ScreenshotsResolved: answeredCount,
	//	ScreenshotsLeft:     totalCount - answeredCount,
	//	UsersUnique:         len(usersMap),
	//	UsersActive:         activityUsers,
	//}
	//
	//return ctx.JSON(
	//	http.StatusOK,
	//	httputils.BuildSuccessResponse(resp),
	//)
}
