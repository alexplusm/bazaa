package controllers

import (
	"github.com/labstack/echo"
)

type StatisticsUserQP struct {
	ExtSystemID ExtSystemIDQP
	TotalOnly   TotalOnlyQP
	GameIDs     GameIDsQP
	Duration    DurationQP
}

func (qp *StatisticsUserQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.TotalOnly = buildTotalOnlyQP(ctx)
	qp.GameIDs = buildGameIDsQP(ctx)
	qp.Duration = buildDurationQP(ctx)
}

// ---

type StatisticsLeaderboardQP struct {
	ExtSystemID ExtSystemIDQP
	Limit       LimitQP
	GameIDs     GameIDsQP
	Duration    DurationQP
}

func (qp *StatisticsLeaderboardQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.Limit = buildLimitQP(ctx)
	qp.GameIDs = buildGameIDsQP(ctx)
	qp.Duration = buildDurationQP(ctx)
}

// ---

type ScreenshotRetrieveQP struct {
	ExtSystemID ExtSystemIDQP
	UserID      UserIDQP
}

func (qp *ScreenshotRetrieveQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.UserID = buildUserIDQP(ctx)
}
