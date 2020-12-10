package controllers

import (
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

// ---

type DurationQP struct {
	From string
	To   string
}

func buildDurationQP(ctx echo.Context) DurationQP {
	return DurationQP{
		From: ctx.QueryParam(consts.FromQPName),
		To:   ctx.QueryParam(consts.ToQPName),
	}
}

// ---

type GameIDsQP struct {
	Value []string
}

func buildGameIDsQP(ctx echo.Context) GameIDsQP {
	gameIDsQP := ctx.QueryParam(consts.GameIDsQPName)
	gameIDs := make([]string, 0, 16)
	if gameIDsQP != "" {
		rowGameIDs := strings.Split(gameIDsQP, ",")
		for _, rowGameID := range rowGameIDs {
			gameIDs = append(gameIDs, strings.Trim(rowGameID, " "))
		}
	}
	return GameIDsQP{Value: gameIDs}
}

// ---

type TotalOnlyQP struct {
	Value bool
}

func buildTotalOnlyQP(ctx echo.Context) TotalOnlyQP {
	totalOnly := ctx.QueryParam(consts.TotalOnlyQPName)
	return TotalOnlyQP{Value: strings.ToLower(totalOnly) == "true"}
}

// ---

type ExtSystemIDQP struct {
	Value string
}

func buildExtSystemIDQP(ctx echo.Context) ExtSystemIDQP {
	qp := ctx.QueryParam(consts.ExtSystemIDQPName)
	return ExtSystemIDQP{Value: qp}
}

// ---

type LimitQP struct {
	Value int
}

func buildLimitQP(ctx echo.Context) LimitQP {
	limitRaw := ctx.QueryParam(consts.LimitQPName)
	limit := 6 // TODO: consts

	if limitRaw != "" {
		intLimit64, err := strconv.ParseInt(limitRaw, 10, 64)
		if err == nil {
			limit = int(intLimit64)
		}
	}

	return LimitQP{Value: limit}
}

// ---

type UserIDQP struct {
	Value string
}

func buildUserIDQP(ctx echo.Context) UserIDQP {
	qp := ctx.QueryParam(consts.UserIDQPName)
	return UserIDQP{Value: qp}
}
