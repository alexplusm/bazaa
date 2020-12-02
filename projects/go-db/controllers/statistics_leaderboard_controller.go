package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsLeaderboardController struct {
	ExtSystemService interfaces.IExtSystemService
	GameService      interfaces.IGameService
	AnswerService    interfaces.IAnswerService
}

func (controller *StatisticsLeaderboardController) GetStatistics(ctx echo.Context) error {
	limit := ctx.QueryParam(consts.LimitQueryParamName)

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
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem not found"),
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

	println("expectedGames: ", expectedGames)

	if len(expectedGames) == 0 {
		fmt.Println("Zero game")
		return nil
	}

	firstGame := expectedGames[0]
	for _, game := range expectedGames {
		if firstGame.StartDate.Before(game.StartDate) {
			firstGame = game
		}
	}

	// TODO: copypast
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
	// TODO: copypast

	var intLimit int
	if limit != "" {
		intLimit64, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			log.Error("todo", err) // TODO
			intLimit = 6
		}
		intLimit = int(intLimit64)
	} else {
		intLimit = 6
	}

	listDAO := make([]dao.AnswerStatLeadDAO, 0, 1024)

	for _, game := range expectedGames {
		// TODO:
		res, err := controller.AnswerService.ABC(game.GameID, fromTime, toTime)
		if err != nil {
			// TODO: LOG
			continue
		}
		fmt.Println("Res: ", res)
		listDAO = append(listDAO, res...)
	}

	// MAP

	userAnswerMap := make(map[string]bo.StatisticsUsersInner)

	for _, listItem := range listDAO {
		_, ok := userAnswerMap[listItem.UserID]
		if !ok {
			userAnswerMap[listItem.UserID] = bo.StatisticsUsersInner{}
		}

		val := userAnswerMap[listItem.UserID]

		if listItem.Value == listItem.UsersAnswer {
			val.RightAnswers++
			//userAnswerMap[listItem.UserID].RightAnswers++
		}
		if listItem.ExpertAnswer == listItem.Value {
			val.MatchWithExpert++
			//userAnswerMap[listItem.UserID].MatchWithExpert++
		}
		val.TotalScreenshots++
		//userAnswerMap[listItem.UserID].TotalScreenshots++

		userAnswerMap[listItem.UserID] = val
	}

	fmt.Println(intLimit)
	//fmt.Println("LIST: ", listDAO)
	fmt.Printf("MAP: %+v\n", userAnswerMap)

	leaders := make([]dto.LeadersDTO, 0, 1024)

	for key, item := range userAnswerMap {
		sts := dto.StatisticsUsersInnerDTO{
			TotalScreenshots: item.TotalScreenshots,
			RightAnswers:     item.RightAnswers,
			MatchWithExpert:  item.MatchWithExpert,
			AverageAccuracy:  item.AverageAccuracy,
		}
		dtooo := dto.LeadersDTO{UserID: key, Statistics: sts}
		leaders = append(leaders, dtooo)
	}

	sort.Slice(leaders, func(i, j int) bool {
		return leaders[i].Statistics.RightAnswers > leaders[j].Statistics.RightAnswers
	})

	leaders = leaders[:intLimit]

	resp := dto.LeadersResponseDTO{Leaders: leaders}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
