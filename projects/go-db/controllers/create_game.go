package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/models"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/errors"
)

// source: https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508

// source: https://godoc.org/github.com/go-playground/validator
// TODO: package lvl variable - yes: move to package lvl file with variable | concurrency - need lock???
// TODO: init func ????
var validate *validator.Validate

type CreateGameController struct {
	Service interfaces.IGameService
}

// CreateGameС create game controller
// "application/json" content-type only - make middleware?
// TODO: rename to "CreateGame"
// TODO: tests
func (controller *CreateGameController) CreateGame(ctx echo.Context) error {
	fmt.Println("CreateGame controller")

	gameRaw := new(dto.CreateGameRequestBody)

	if err := ctx.Bind(gameRaw); err != nil {
		fmt.Printf("CreateGame controller: %v\n", err)
	}

	validate = validator.New()

	//g := new(dto.Game)
	g := new(models.GameModel)
	err := g.CreateGame(*gameRaw, validate)
	if err != nil {
		ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
		return fmt.Errorf("Create Game controller: %v", err)
	}

	fmt.Printf("GAME: %+v\n", g)

	controller.Service.CreateGame()

	return nil
}
