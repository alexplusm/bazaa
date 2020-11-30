package dto

import (
	"strings"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

type DurationQueryParams struct {
	From string
	To   string
}

type StatisticsUserQueryParams struct {
	ExtSystemID string
	TotalOnly   bool
	GameIDs     []string
	Duration    DurationQueryParams
}

func (qp *DurationQueryParams) FromCTX(ctx echo.Context) {
	qp.From = ctx.QueryParam(consts.FromQueryParamName)
	qp.To = ctx.QueryParam(consts.ToQueryParamName)
}

func (qp *StatisticsUserQueryParams) FromCTX(ctx echo.Context) {
	gameIDs := ctx.QueryParam(consts.GameIDsQueryParamName)
	totalOnly := ctx.QueryParam(consts.TotalOnlyQueryParamName)
	duration := DurationQueryParams{}
	duration.FromCTX(ctx)

	qp.ExtSystemID = ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	qp.TotalOnly = strings.ToLower(totalOnly) == "true"
	qp.Duration = duration

	if gameIDs == "" {
		qp.GameIDs = make([]string, 0, 0)
	} else {
		qp.GameIDs = strings.Split(gameIDs, ",")
	}
}
