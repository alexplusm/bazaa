package controllers

import (
	"fmt"
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

	fmt.Println(userID)
	fmt.Printf("Query Params: %+v\n", qp)

	// TODO: обобщенная функция по обработке ошибок или не существования сущности
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

	userExist, err := controller.UserService.UserExist(userID)
	if err != nil {
		log.Error("get user statistics controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !userExist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("user not found"),
		)
	}

	games, err := controller.GameService.GetGames(qp.ExtSystemID.Value)
	if len(games) == 0 {
		// TODO:discuss: что делать в этом случае
		// * empty statistics
		// * error: game not found
	}

	games = controller.GameService.FilterGames(qp.GameIDs.Value, games)

	expectedGames := games

	if len(games) == 0 {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}

	earliestGame := controller.GameService.GetEarliestGame(games)

	// TODO:discuss: если ошибка в парсинге даты?
	// * оповещать пользователя
	// * использовать дефолтные значения
	from, to := controller.DurationService.GetDurationByGame(
		qp.Duration.From, qp.Duration.To, earliestGame,
	)

	stats, err := controller.AnswerService.GetUserStatistics(userID, qp.TotalOnly.Value, expectedGames, from, to)
	if err != nil {
		log.Error("get user statistics controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	var resp interface{}

	if qp.TotalOnly.Value {
		resp = bo.StatsToTotalOnlyDTO(stats)
	} else {
		resp = bo.StatsToDTO(stats)
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
