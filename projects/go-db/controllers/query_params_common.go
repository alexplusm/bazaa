package controllers

import (
	"strconv"
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

type LimitQueryParam struct {
	Value int
}

func buildLimitQueryParam(ctx echo.Context) LimitQueryParam {
	limitRaw := ctx.QueryParam(consts.LimitQueryParamName)
	limit := 6 // TODO: consts

	if limitRaw != "" {
		intLimit64, err := strconv.ParseInt(limitRaw, 10, 64)
		if err == nil {
			limit = int(intLimit64)
		}
	}

	return LimitQueryParam{Value: limit}
}

// ---

type UserIDQP struct {
	Value string
}

func buildUserIDQP(ctx echo.Context) UserIDQP {
	qp := ctx.QueryParam(consts.UserIDQueryParamName)
	return UserIDQP{Value: qp}
}
