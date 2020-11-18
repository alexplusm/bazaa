package bo

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type GameBO struct {
	Name       string    `validate:"required"`
	AnswerType int       `validate:"gte=1,lte=4"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time `validate:"required,gtcsfield=StartDate"`
	Question   string    `validate:"required"`
	Options    string
}

func (g *GameBO) CreateGame(src dto.CreateGameRequestBody, validate *validator.Validate) error {
	startDate, err := strconv.ParseInt(src.StartDate, 10, 64)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}
	endDate, err := strconv.ParseInt(src.EndDate, 10, 64)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}

	g.StartDate = time.Unix(startDate, 0)
	g.EndDate = time.Unix(endDate, 0)

	upperBound := time.Now().AddDate(5, 0, 0)
	if g.StartDate.After(upperBound) || g.EndDate.After(upperBound) {
		return fmt.Errorf("create game: StartDate or EndDate too far dates")
	}

	// TODO: нельзя на прошлую дату создать игру: добавить валидацию

	g.Name = src.Name
	g.AnswerType = src.AnswerType
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

func (g *GameBO) validate() error {
	if g.AnswerType == consts.CategoricalAnswerType {
		options := strings.Split(g.Options, ",")

		if len(options) < 2 {
			return fmt.Errorf("required 2 options or more")
		}
		for _, s := range options {
			if s == "" {
				return fmt.Errorf("required option value")
			}
		}
	}

	return nil
}

func (g *GameBO) NotStarted() bool {
	return time.Now().Before(g.StartDate)
}
