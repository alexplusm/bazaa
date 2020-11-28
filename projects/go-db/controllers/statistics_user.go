package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

type StatisticsUserController struct {
}

func (controller StatisticsUserController) GetStatistics(ctx echo.Context) error {
	// TODO: params in urls -> consts
	userID := ctx.Param("user-id")
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	totalOnly := ctx.QueryParam(consts.TotalOnlyQueryParamName)
	gameIDs := ctx.QueryParam(consts.GameIDsQueryParamName)
	from := ctx.QueryParam(consts.FromQueryParamName)
	to := ctx.QueryParam(consts.ToQueryParamName)

	fmt.Println(userID, extSystemID, totalOnly, gameIDs, from, to)

	return nil
}
