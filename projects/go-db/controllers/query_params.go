package controllers

import (
	"github.com/labstack/echo"
)

type StatisticsUserQueryParams struct {
	ExtSystemID ExtSystemIDQueryParam
	TotalOnly   TotalOnlyQueryParam
	GameIDs     GameIDsQueryParam
	Duration    DurationQueryParams
}

func (qp *StatisticsUserQueryParams) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.TotalOnly = buildTotalOnlyQueryParam(ctx)
	qp.GameIDs = buildGameIDsQueryParam(ctx)
	qp.Duration = buildDurationQueryParams(ctx)
}

// ---

type StatisticsLeaderBoardQueryParams struct {
	ExtSystemID ExtSystemIDQueryParam
	Limit       LimitQueryParam
	GameIDs     GameIDsQueryParam
	Duration    DurationQueryParams
}

func (qp *StatisticsLeaderBoardQueryParams) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.Limit = buildLimitQueryParam(ctx)
	qp.GameIDs = buildGameIDsQueryParam(ctx)
	qp.Duration = buildDurationQueryParams(ctx)
}
