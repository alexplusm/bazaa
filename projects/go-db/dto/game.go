package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// source: https://godoc.org/github.com/go-playground/validator
// package lvl variable - yes: move to package lvl file with variable | concurrency - need lock???
// init func ????

// TODO: rename
type CreateGameRequestBody struct {
	Name       string `json:"name"`
	AnswerType int    `json:"answer_type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Question   string `json:"question"`
	Options    string `json:"options"`
}

// TODO: rename
type Game struct {
	Name       string    `validate:"required"`
	AnswerType int       `validate:"gte=0,lte=2"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time `validate:"required,gtcsfield=StartDate"`
	Question   string    `validate:"required"`
	Options    string    `validate:"required"`
}

// TODO: rename val
func (g *Game) CreateGame(src CreateGameRequestBody, val *validator.Validate) error {
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

	if err := val.Struct(g); err != nil {
		return fmt.Errorf("CreateGame validation: %v", err)
	}
	if err := g.validate(); err != nil {
		return fmt.Errorf("CreateGame validation: %v", err)
	}

	fmt.Printf("Created game: %+v\n", g) // TODO: log game creation
	return nil
}

func (g *Game) validate() error {
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
