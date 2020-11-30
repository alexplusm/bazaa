package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type StatisticsUserController struct {
	ExtSystemService interfaces.IExtSystemService
	AnswerService    interfaces.IAnswerService
}

func (controller StatisticsUserController) GetStatistics(ctx echo.Context) error {
	// TODO: params in urls -> consts
	userID := ctx.Param("user-id")

	qp := dto.StatisticsUserQueryParams{}
	qp.FromCTX(ctx)

	fmt.Printf("Params: %+v\n", qp)

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

	fmt.Println(userID)

	// TODO: VALIDATE query params
	// TODO: add column: answer_date | for filtering

	//controller.AnswerService.GetUserStatistics(userID, totalOnly, gameIDs, from, to)

	//fmt.Println("extSystemID: ", extSystemID)
	//
	//fmt.Println(userID, totalOnly, gameIDs, from, to)

	return nil
}
