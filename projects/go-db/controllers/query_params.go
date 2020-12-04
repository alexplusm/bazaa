package controllers

import (
	"github.com/labstack/echo"
)

type StatisticUserQP struct {
	ExtSystemID ExtSystemIDQP
	TotalOnly   TotalOnlyQP
	GameIDs     GameIDsQP
	Duration    DurationQP
}

func (qp *StatisticUserQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.TotalOnly = buildTotalOnlyQP(ctx)
	qp.GameIDs = buildGameIDsQP(ctx)
	qp.Duration = buildDurationQP(ctx)
}

// ---

type StatisticLeaderboardQP struct {
	ExtSystemID ExtSystemIDQP
	Limit       LimitQP
	GameIDs     GameIDsQP
	Duration    DurationQP
}

func (qp *StatisticLeaderboardQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.Limit = buildLimitQP(ctx)
	qp.GameIDs = buildGameIDsQP(ctx)
	qp.Duration = buildDurationQP(ctx)
}

// ---

type StatisticGameQP struct {
	ExtSystemID ExtSystemIDQP
	GameIDs     GameIDsQP
}

func (qp *StatisticGameQP) fromCtx(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQP(ctx)
	qp.GameIDs = buildGameIDsQP(ctx)
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
