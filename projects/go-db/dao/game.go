package dao

import (
	"github.com/Alexplusm/bazaa/projects/go-db/domain"
)

type GameDAO struct {
	Name       string
	AnswerType int
	StartDate  int64
	EndDate    int64
	Question   string
	Options    string
}

func (g *GameDAO) FromBO(bo domain.GameBO) {
	g.Name = bo.Name
	g.AnswerType = bo.AnswerType

	g.StartDate = bo.StartDate.Unix()
	g.EndDate = bo.EndDate.Unix()

	g.Question = bo.Question
	g.Options = bo.Options
}
