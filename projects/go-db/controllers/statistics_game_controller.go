package controllers

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsGameController struct {
	GameService       interfaces.IGameService
	ExtSystemService  interfaces.IExtSystemService
	ScreenshotService interfaces.IScreenshotService
}

func (controller *StatisticsGameController) GetStatistics(ctx echo.Context) error {
	// TODO: qp
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	//gameIDs := ctx.QueryParam(consts.GameIDsQueryParamName)

	exist, err := controller.ExtSystemService.ExtSystemExist(extSystemID)
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

	qp := dto.StatisticsUserQueryParams{}
	qp.FromCTX(ctx)

	fmt.Println("AZAZAZAZAZAZAZAZZAZAZA", qp.ExtSystemID)

	games, err := controller.GameService.GetGames(qp.ExtSystemID)
	fmt.Println("Error: ", err)
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

	fmt.Printf("GAMES: %+v\n", expectedGames)

	//filteredGamesIDs := make([]string, 0, len(expectedGames))
	//for _, g := range expectedGames {
	//	filteredGamesIDs = append(filteredGamesIDs, g.GameID)
	//}

	count := 0

	for _, g := range expectedGames {
		c, _ := controller.ScreenshotService.ScreenshotCountByGame(g.GameID)
		count += c
	}

	fmt.Println("COUNT: ", count)

	// get Screenshots

	resp := dto.GameStatsDTO{}

	return ctx.JSON(
		http.StatusOK,
		httputils.BuildSuccessResponse(resp),
	)
}
