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
	GameID      string
	ExtSystemID string    `validate:"required"`
	Name        string    `validate:"required"`
	AnswerType  int       `validate:"gte=1,lte=4"`
	StartDate   time.Time `validate:"required"`
	EndDate     time.Time `validate:"required,gtcsfield=StartDate"`
	Question    string    `validate:"required"`
	Options     string
}

func (g *GameBO) FromDTO(src dto.CreateGameRequestBody, validate *validator.Validate) error {
	startDate, err := strconv.ParseInt(src.StartDate, 10, 64)
	if err != nil {
		return fmt.Errorf("game: from dto: %v", err)
	}
	endDate, err := strconv.ParseInt(src.EndDate, 10, 64)
	if err != nil {
		return fmt.Errorf("game: from dto: %v", err)
	}

	// TODO: use FromTimestampToTime

	g.StartDate = time.Unix(startDate, 0)
	g.EndDate = time.Unix(endDate, 0)

	now := time.Now()
	if g.StartDate.Before(now) || g.EndDate.Before(now) {
		return fmt.Errorf("game: from dto: StartDate or EndDate in the past")
	}
	upperBound := now.AddDate(5, 0, 0)
	if g.StartDate.After(upperBound) || g.EndDate.After(upperBound) {
		return fmt.Errorf("game: from dto: StartDate or EndDate too far dates")
	}

	g.ExtSystemID = src.ExtSystemID
	g.Name = src.Name
	g.AnswerType = src.AnswerType
	g.Question = src.Question
	g.Options = src.Options

	if err := validate.Struct(g); err != nil {
		return fmt.Errorf("game: from dto: %v", err)
	}
	if err := g.validate(); err != nil {
		return fmt.Errorf("game: from dto: %v", err)
	}

	return nil
}

func (g *GameBO) ToListItemDTO() dto.GameItemResponseBody {
	gameDTO := new(dto.GameItemResponseBody)
	gameDTO.GameID = g.GameID
	gameDTO.Name = g.Name
	gameDTO.Status = g.Status()
	gameDTO.From = strconv.FormatInt(g.StartDate.Unix(), 10)
	gameDTO.To = strconv.FormatInt(g.EndDate.Unix(), 10)

	return *gameDTO
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

func (g *GameBO) Status() string {
	// TODO: statuses in consts
	switch {
	case g.InProcess():
		return "inProcess"
	case g.Finished():
		return "finished"
	}
	return "notStarted"
}

func (g *GameBO) NotStarted() bool {
	return time.Now().Before(g.StartDate)
}

func (g *GameBO) InProcess() bool {
	now := time.Now()
	return now.After(g.StartDate) && now.Before(g.EndDate)
}

func (g *GameBO) Finished() bool {
	return time.Now().After(g.EndDate)
}
