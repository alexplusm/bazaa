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

type StatisticsGameController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *StatisticsGameController) GetStatistics(ctx echo.Context) error {
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	gameIDs := ctx.QueryParam(consts.GameIDsQueryParamName)

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

	fmt.Println(gameIDs)

	return nil
}
