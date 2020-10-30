package controllers

import (
	"fmt"
	"time"

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

type timestamp time.Time

type game struct {
	Name       string
	AnswerType int
	StartDate  timestamp
	EndDate    timestamp
	Question   string
	Options    string
}

func (g *game) createGame(src createGameRequestBody) error {
	startDate, err := time.Parse(time.RFC3339, src.StartDate)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}
	endDate, err := time.Parse(time.RFC3339, src.EndDate)
	if err != nil {
		return fmt.Errorf("CreateGame: %v", err)
	}

	g.StartDate = timestamp(startDate)
	g.EndDate = timestamp(endDate)
	g.Name = src.Name
	g.Question = src.Question
	g.Options = src.Options

	// todo: validate

	fmt.Printf("startDate %+v | %+v\n", g, err)

	// todo: log game creation

	return nil
}

// func (t *Timestamp) UnmarshalParam(src string) error {
// 	ts, err := time.Parse(time.RFC3339, src)
// 	*t = Timestamp(ts)
// 	return err
// }

// source: https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508

// CreateGame create game controller
// application/json only - make middleware
func CreateGame(p *pgxpool.Pool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		gameRaw := new(createGameRequestBody)

		if err := ctx.Bind(gameRaw); err != nil {
			fmt.Printf("CreateGame controller: %v\n", err)
		}

		fmt.Printf("Game Raw val %+v\n", gameRaw)

		g := new(game)

		g.createGame(*gameRaw)

		// TODO: validate Game: if error -> bad request
		// (check on default values all fields)
		// https://echo.labstack.com/guide/request | search "Validate Data"

		return nil
	}
}

// func formDataWithFiles(f c.Form) {}

func test(ctx echo.Context) {

	game := new(createGameRequestBody)

	if err := ctx.Bind(game); err != nil {
		fmt.Printf("CreateGame controller: %v\n", err)
	}

	fmt.Printf("Game val %+v\n", game)

}
