package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/Alexplusm/bazaa/projects/go-db/dto"
)

type GameModel struct {
	Name       string    `validate:"required"`
	AnswerType int       `validate:"gte=1,lte=4"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time `validate:"required,gtcsfield=StartDate"`
	Question   string    `validate:"required"`
	Options    string    `validate:"required"`
}

func (g *GameModel) CreateGame(src dto.CreateGameRequestBody, validate *validator.Validate) error {
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

	fmt.Printf("Created game: %+v\n", g) // TODO:log: game creation
	return nil
}

func (g *GameModel) validate() error {
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
