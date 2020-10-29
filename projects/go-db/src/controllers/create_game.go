package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
)

type schedule struct {
	ID string `json:"id"`
}

type gameRequestBody struct {
	AnswerType string     `json:"answer_type" form:"answer_type"`
	StartDate  string     `json:"start_date" form:"start_date"`
	EndDate    string     `json:"end_date" form:"end_date"`
	Shedules   []schedule `json:"schedules" form:"schedules"`
}

// multipart/form-data
// application/json

// source: https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508

// CreateGame create game controller
func CreateGame(p *pgxpool.Pool) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		fmt.Println("Request accepted!")

		// TODO: may be stdlib func this same functional exsists?
		contentTypeRaw := ctx.Request().Header["Content-Type"][0]
		contentType := strings.Split(contentTypeRaw, ";")[0]

		// temp: testing
		if contentType == "application/json" {
			test(ctx)
		}
		// temp: testing end

		// TODO: process type "application/json" (use c.Bind() echo func)
		if contentType != "multipart/form-data" {
			errMsg := errors.GetBadRequestErrorResponseJSONStr()

			// ctx.JSON() // TODO
			ctx.String(http.StatusOK, errMsg)
			return nil
		}

		// bisness logic
		form, err := ctx.MultipartForm()
		if err != nil {
			return fmt.Errorf("CreateGame controller: %v", err) // TODO: log this error
		}

		archives := form.File["archives"]

		if len(archives) == 0 {
			// todo: schedules branch logic
		} else {
			// todo: files branch logic
		}

		// --- testing
		v := form.Value
		fmt.Printf("FormValue: %+v\n", v)
		fmt.Println(v["id"])
		test(ctx)

		// TODO: validate Game: if error -> bad request
		// (check on default values all fields)
		// https://echo.labstack.com/guide/request | search "Validate Data"

		return nil
	}
}

// func formDataWithFiles(f c.Form) {}

func test(ctx echo.Context) {

	game := new(gameRequestBody)

	if err := ctx.Bind(game); err != nil {
		fmt.Printf("CreateGame controller: %v\n", err)
	}

	fmt.Printf("Game val %+v\n", game)

}
