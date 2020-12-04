package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticLeaderboardController struct {
	ExtSystemService   interfaces.IExtSystemService
	GameService        interfaces.IGameService
	DurationService    interfaces.IDurationService
	LeaderboardService interfaces.ILeaderboardService
}

func (controller *StatisticLeaderboardController) GetStatistics(ctx echo.Context) error {
	qp := StatisticLeaderboardQP{}
	qp.fromCtx(ctx)

	exist, err := controller.ExtSystemService.ExtSystemExist(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("leaderboard statistic controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem not found"),
		)
	}

	games, err := controller.GameService.GetGames(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("leaderboard statistic controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	games = controller.GameService.FilterGames(qp.GameIDs.Value, games)
	if len(games) == 0 {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}

	earliestGame := controller.GameService.GetEarliestGame(games)

	from, to, err := controller.DurationService.GetDurationByGame(
		qp.Duration.From, qp.Duration.To, earliestGame,
	)
	if err != nil {
		log.Error("user statistic controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("error while date parsing"),
		)
	}

	gameIDs := make([]string, 0, len(games))
	for _, game := range games {
		gameIDs = append(gameIDs, game.GameID)
	}

	resp := controller.LeaderboardService.GetLeaderboard(gameIDs, from, to, qp.Limit.Value)

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
