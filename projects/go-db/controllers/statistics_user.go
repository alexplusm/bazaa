package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsUserController struct {
	ExtSystemService interfaces.IExtSystemService
	AnswerService    interfaces.IAnswerService
}

func (controller StatisticsUserController) GetStatistics(ctx echo.Context) error {
	// TODO: params in urls -> consts
	userID := ctx.Param("user-id")
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	totalOnly := ctx.QueryParam(consts.TotalOnlyQueryParamName)
	gameIDs := ctx.QueryParam(consts.GameIDsQueryParamName)
	from := ctx.QueryParam(consts.FromQueryParamName)
	to := ctx.QueryParam(consts.ToQueryParamName)

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

	// TODO: add column: answer_date | for filtering

	fmt.Println("extSystemID: ", extSystemID)

	fmt.Println(userID, totalOnly, gameIDs, from, to)

	return nil
}
