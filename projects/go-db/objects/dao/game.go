package dao

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type GameDAO struct {
	GameID      string
	ExtSystemID string
	Name        string
	AnswerType  int
	StartDate   int64
	EndDate     int64
	Question    string
	Options     string
}

func (g *GameDAO) FromBO(bo bo.GameBO) {
	g.ExtSystemID = bo.ExtSystemID
	g.Name = bo.Name
	g.AnswerType = bo.AnswerType
	g.StartDate = bo.StartDate.Unix()
	g.EndDate = bo.EndDate.Unix()
	g.Question = bo.Question
	g.Options = bo.Options
}

func (g *GameDAO) ToBO() bo.GameBO {
	gameBO := bo.GameBO{}

	gameBO.GameID = g.GameID
	gameBO.ExtSystemID = g.ExtSystemID
	gameBO.Name = g.Name
	gameBO.StartDate = time.Unix(g.StartDate, 0)
	gameBO.EndDate = time.Unix(g.EndDate, 0)
	gameBO.AnswerType = g.AnswerType
	gameBO.Question = g.Question
	gameBO.Options = g.Options // TODO: not required

	return gameBO
}
