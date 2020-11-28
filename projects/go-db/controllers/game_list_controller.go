package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameListController struct {
	GameService      interfaces.IGameService
	ExtSystemService interfaces.IExtSystemService
}

func (controller *GameListController) GetGames(ctx echo.Context) error {
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)

	exist, err := controller.ExtSystemService.ExtSystemExist(extSystemID)
	if err != nil {
		log.Error("get games controller: ", err)
		// TODO: return bad request or internal system error?
	}

	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem not found"),
		)
	}

	fmt.Println("extSystemID: ", extSystemID)

	gamesBO, err := controller.GameService.GetGames(extSystemID)
	if err != nil {
		log.Error("get games controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	gamesDTO := make([]dto.GameItemResponseBody, 0, len(gamesBO))
	for _, game := range gamesBO {
		gamesDTO = append(gamesDTO, game.ToListItemDTO())
	}
	body := dto.GameListResponseBody{Games: gamesDTO}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(body))
}
