package controllers

import (
	"fmt"
	"net/http"

	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/errors"
	"github.com/Alexplusm/bazaa/projects/go-db/src/utils/middlewares"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
)

// TODO: move
const (
	// FormDataType const
	FormDataType = "multipart/form-data"
	// ApplicationJSON const
	ApplicationJSON = "application/json"
)

// UpdateGame update game controller
func UpdateGame(p *pgxpool.Pool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("GAME ID", ctx.Param("game-id"))
		switch middlewares.ParseContentType(ctx) {
		case FormDataType:
			archives, err := ctx.FormFile("archives")
			if err != nil {
				// TODO: log error
				ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
				return fmt.Errorf("UpdateGame controller: %v", err)
			}
			fmt.Println(archives)
			fmt.Println("FORM DATA TYPE")
		case ApplicationJSON:
			fmt.Println("APPLICATION JSON")
		default:
			ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
		}
		return nil
	}
}
