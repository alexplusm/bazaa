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
}

func (controller *StatisticGameController) GetStatistics(ctx echo.Context) error {
	qp := StatisticGameQP{}
	qp.fromCtx(ctx)

	if qp.ExtSystemID.Value == "" {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem required"),
		)
	}

	exist, err := controller.ExtSystemService.Exist(qp.ExtSystemID.Value)
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

	games, err := controller.GameService.List(qp.ExtSystemID.Value)
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
}
