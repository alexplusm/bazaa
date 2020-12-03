package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsUserController struct {
	GameService      interfaces.IGameService
	ExtSystemService interfaces.IExtSystemService
	AnswerService    interfaces.IAnswerService
	UserService      interfaces.IUserService
	DurationService  interfaces.IDurationService
}

func (controller StatisticsUserController) GetStatistics(ctx echo.Context) error {
	userID := ctx.Param(consts.UserIDUrlParam)
	qp := StatisticsUserQueryParams{}
	qp.fromCtx(ctx)

	exist, err := controller.ExtSystemService.ExtSystemExist(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("get user statistics controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("extSystem not found"),
		)
	}

	exist, err = controller.UserService.UserExist(userID)
	if err != nil {
		log.Error("get user statistics controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("user not found"),
		)
	}

	games, err := controller.GameService.GetGames(qp.ExtSystemID.Value)
	if err != nil {
		log.Error("get user statistics controller: ", err)
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

	// TODO:discuss: если ошибка в парсинге даты?
	// 1) оповещать пользователя
	// 2) использовать дефолтные значения
	from, to := controller.DurationService.GetDurationByGame(
		qp.Duration.From, qp.Duration.To, earliestGame,
	)

	gameIDs := make([]string, 0, len(games))
	for _, game := range games {
		gameIDs = append(gameIDs, game.GameID)
	}

	// TODO: statistics service
	statistics, err := controller.AnswerService.GetUserStatistics(userID, gameIDs, from, to)
	if err != nil {
		log.Error("get user statistics controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	var resp interface{}

	if qp.TotalOnly.Value {
		resp = bo.StatisticAnswersDateSlicedBOToDTOTotalOnly(statistics) // TODO
	} else {
		resp = bo.StatisticAnswersDateSlicedBOToDTO(statistics) // TODO
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
