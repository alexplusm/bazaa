package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsUserController struct {
	GameService      interfaces.IGameService
	ExtSystemService interfaces.IExtSystemService
	AnswerService    interfaces.IAnswerService
	UserService      interfaces.IUserService
}

func (controller StatisticsUserController) GetStatistics(ctx echo.Context) error {
	// TODO: params in urls -> consts
	userID := ctx.Param("user-id")
	// получить самую раннюю игру
	fmt.Println(userID)

	qp := dto.StatisticsUserQueryParams{}
	qp.FromCTX(ctx)

	exist, err := controller.ExtSystemService.ExtSystemExist(qp.ExtSystemID)
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

	if len(expectedGames) == 0 {
		// TODO: log
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}

	firstGame := expectedGames[0]
	for _, game := range expectedGames {
		if firstGame.StartDate.Before(game.StartDate) {
			firstGame = game
		}
	}

	// Validation
	year, month, day := time.Now().Date()
	var fromTime = firstGame.StartDate
	var toTime = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	if qp.Duration.From != "" {
		fromm, err := utils.FromTimestampToTime(qp.Duration.From)
		if err != nil {
			// TODO: log
		} else {
			fromTime = fromm
		}
	}
	if qp.Duration.To != "" {
		to_o, err := utils.FromTimestampToTime(qp.Duration.To)
		if err != nil {
			// TODO: log
		} else {
			toTime = to_o
		}
	}

	stats, err := controller.AnswerService.GetUserStatistics(userID, qp.TotalOnly, expectedGames, fromTime, toTime)
	if err != nil {
		// TODO: LOG ERROR
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	if qp.TotalOnly {
		resp := bo.StatsToTotalOnlyDTO(stats)
		return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
	} else {
		resp := bo.StatsToDTO(stats)
		return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
	}
}
