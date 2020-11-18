package dao

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type GameDAO struct {
	Name       string
	AnswerType int
	StartDate  int64
	EndDate    int64
	Question   string
	Options    string
}

func (g *GameDAO) FromBO(bo bo.GameBO) {
	g.Name = bo.Name
	g.AnswerType = bo.AnswerType
	g.StartDate = bo.StartDate.Unix()
	g.EndDate = bo.EndDate.Unix()
	g.Question = bo.Question
	g.Options = bo.Options
}

func (g *GameDAO) ToBO() bo.GameBO {
	gameBO := new(bo.GameBO)

	gameBO.Name = g.Name
	gameBO.StartDate = time.Unix(g.StartDate, 0)
	gameBO.EndDate = time.Unix(g.EndDate, 0)
	gameBO.AnswerType = g.AnswerType
	gameBO.Question = g.Question
	gameBO.Options = g.Options

	return *gameBO
}
