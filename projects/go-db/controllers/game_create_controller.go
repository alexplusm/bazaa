package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

// source: https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508

type GameCreateController struct {
	Service interfaces.IGameService
}

func (controller *GameCreateController) CreateGame(ctx echo.Context) error {
	gameRaw := new(dto.CreateGameRequestBody)

	if err := ctx.Bind(gameRaw); err != nil {
		fmt.Printf("CreateGame controller: %v\n", err)
	}

	fmt.Printf("GameRaw %+v\n", gameRaw)

	game := new(bo.GameBO)
	if err := game.CreateGame(*gameRaw, validate); err != nil {
		ctx.String(http.StatusOK, httputils.GetBadRequestErrorResponseJSONStr())
		return fmt.Errorf("create game controller: %v", err)
	}

	gameID, err := controller.Service.CreateGame(*game)
	if err != nil {
		return fmt.Errorf("create game controller: %v", err)
	}

	ctx.String(http.StatusOK, "gameID: "+gameID) // todo: response generator

	return nil
}
