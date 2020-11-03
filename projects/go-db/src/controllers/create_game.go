package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/errors"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
)

type createGameRequestBody struct {
	Name       string `json:"name"`
	AnswerType int    `json:"answer_type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Question   string `json:"question"`
	Options    string `json:"options"`
}

type game struct {
	Name       string    `validate:"required"`
	AnswerType int       `validate:"gte=0,lte=2"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time `validate:"required,gtcsfield=StartDate"`
	Question   string    `validate:"required"`
	Options    string    `validate:"required"`
}

// source: https://godoc.org/github.com/go-playground/validator
// package lvl variable - yes: move to package lvl file with variable | concurrency - need lock???
// init func ????
var validate *validator.Validate

// to model?
func (g *game) createGame(src createGameRequestBody) error {
	validate = validator.New()

	startDate, err := time.Parse(time.RFC3339, src.StartDate)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}
	endDate, err := time.Parse(time.RFC3339, src.EndDate)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}

	g.StartDate = startDate
	g.AnswerType = src.AnswerType
	g.EndDate = endDate
	g.Name = src.Name
	g.Question = src.Question
	g.Options = src.Options

	if err := validate.Struct(g); err != nil {
		return fmt.Errorf("CreateGame validation: %v", err)
	}
	if err := g.validate(); err != nil {
		return fmt.Errorf("CreateGame validation: %v", err)
	}

	fmt.Printf("Created game: %+v\n", g) // TODO: log game creation
	return nil
}

func (g *game) validate() error {
	options := strings.Split(g.Options, ",")

	if len(options) < 2 {
		return fmt.Errorf("required 2 options or more")
	}
	for _, s := range options {
		if s == "" {
			return fmt.Errorf("required option value")
		}
	}

	return nil
}

// source: https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508

// CreateGame create game controller
// "application/json" content-type only - make middleware?
func CreateGame(p *pgxpool.Pool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		gameRaw := new(createGameRequestBody)

		if err := ctx.Bind(gameRaw); err != nil {
			fmt.Printf("CreateGame controller: %v\n", err)
		}

		g := new(game)
		err := g.createGame(*gameRaw)
		if err != nil {
			ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
			return fmt.Errorf("Create Game controller: %v", err)
		}

		// pgx -> create game in DB
		// proccess error

		return nil
	}
}
