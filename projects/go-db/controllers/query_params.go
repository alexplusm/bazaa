package controllers

import (
	"github.com/labstack/echo"
)

type StatisticsUserQP struct {
	ExtSystemID ExtSystemIDQueryParam
	TotalOnly   TotalOnlyQueryParam
	GameIDs     GameIDsQueryParam
	Duration    DurationQueryParams
}

func (qp *StatisticsUserQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.TotalOnly = buildTotalOnlyQueryParam(ctx)
	qp.GameIDs = buildGameIDsQueryParam(ctx)
	qp.Duration = buildDurationQueryParams(ctx)
}

// ---

type StatisticsLeaderboardQP struct {
	ExtSystemID ExtSystemIDQueryParam
	Limit       LimitQueryParam
	GameIDs     GameIDsQueryParam
	Duration    DurationQueryParams
}

func (qp *StatisticsLeaderboardQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.Limit = buildLimitQueryParam(ctx)
	qp.GameIDs = buildGameIDsQueryParam(ctx)
	qp.Duration = buildDurationQueryParams(ctx)
}

// ---

type ScreenshotRetrieveQP struct {
	ExtSystemID ExtSystemIDQueryParam
	UserID      UserIDQP
}

func (qp *ScreenshotRetrieveQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.UserID = buildUserIDQP(ctx)
}
