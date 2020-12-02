package controllers

import (
	"strings"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

// ---

type DurationQueryParams struct {
	From string
	To   string
}

func buildDurationQueryParams(ctx echo.Context) DurationQueryParams {
	return DurationQueryParams{
		From: ctx.QueryParam(consts.FromQueryParamName),
		To:   ctx.QueryParam(consts.ToQueryParamName),
	}
}

// ---

type GameIDsQueryParam struct {
	Value []string
}

func buildGameIDsQueryParam(ctx echo.Context) GameIDsQueryParam {
	gameIDsQP := ctx.QueryParam(consts.GameIDsQueryParamName)
	gameIDs := make([]string, 0, 16)
	if gameIDsQP != "" {
		rowGameIDs := strings.Split(gameIDsQP, ",")
		for _, rowGameID := range rowGameIDs {
			gameIDs = append(gameIDs, strings.Trim(rowGameID, " "))
		}
	}
	return GameIDsQueryParam{Value: gameIDs}
}

// TODO: filter with bo.Games
//func (qp *GameIDsQueryParam) filter() {}

// ---

type TotalOnlyQueryParam struct {
	Value bool
}

func buildTotalOnlyQueryParam(ctx echo.Context) TotalOnlyQueryParam {
	totalOnly := ctx.QueryParam(consts.TotalOnlyQueryParamName)
	return TotalOnlyQueryParam{Value: strings.ToLower(totalOnly) == "true"}
}

// ---

type ExtSystemIDQueryParam struct {
	Value string
}

func buildExtSystemIDQueryParam(ctx echo.Context) ExtSystemIDQueryParam {
	qp := ctx.QueryParam(consts.ExtSystemIDQueryParamName)
	return ExtSystemIDQueryParam{Value: qp}
}

// ---

type StatisticsUserQueryParams struct {
	ExtSystemID ExtSystemIDQueryParam
	TotalOnly   TotalOnlyQueryParam
	GameIDs     GameIDsQueryParam
	Duration    DurationQueryParams
}

func (qp *StatisticsUserQueryParams) FromCTX(ctx echo.Context) {
	qp.ExtSystemID = buildExtSystemIDQueryParam(ctx)
	qp.TotalOnly = buildTotalOnlyQueryParam(ctx)
	qp.GameIDs = buildGameIDsQueryParam(ctx)
	qp.Duration = buildDurationQueryParams(ctx)
}
