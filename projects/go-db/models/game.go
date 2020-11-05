package models

import (
	"fmt"
	"strconv"
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
	startDate, err := strconv.ParseInt(src.StartDate, 10, 64)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}
	endDate, err := strconv.ParseInt(src.EndDate, 10, 64)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}

	g.StartDate = time.Unix(startDate, 0)
	g.AnswerType = src.AnswerType
	g.EndDate = time.Unix(endDate, 0)
	g.Name = src.Name
	g.Question = src.Question
	g.Options = src.Options

	// TODO: test case with AnswerType != 2 (не категориальный тип)
	// то что будет с Options?
	// должна быть пустая строка | хотя нужен nil -> чтобы в базе был NULL
	// (как вариант: 2 sql insertStatement - с options и без)

	// TODO: case: JavaScript генерирует timestamp в миллисекундах
	// чтобы обработать этот кейс - запрещаем создавать игры с датой, отличной от time.Now() + 10 year
	// TODO: запилить эту валидацию и описать "INFO" по этому поводу

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
	// TODO: use ENUM !!!
	if g.AnswerType == 2 {
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
