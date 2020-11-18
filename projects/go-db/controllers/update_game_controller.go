package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type UpdateGameController struct {
	Service interfaces.IUpdateGameService
}

func (controller *UpdateGameController) UpdateGame(ctx echo.Context) error {
	gameID := ctx.Param("game-id")

	fmt.Println("UpdateGameController: GameID:", gameID)

	switch httputils.ParseContentType(ctx) {
	case consts.FormDataContentType:
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusOK, httputils.GetBadRequestErrorResponseJSONStr())
			return fmt.Errorf("update game controller: %v", err)
		}

		game, err := controller.Service.GetGame(gameID)
		if err != nil {
			// TODO: ctx.String: return game NOT found
			return fmt.Errorf("update game controller: %v", err)
		}

		if !game.NotStarted() {
			fmt.Println("Game NOT started")
			// TODO: ctx.String: return game not started
			return nil
		}

		archives := form.File["archives"]

		err = controller.Service.AttachZipArchiveToGame(gameID, archives)
		if err != nil {
			return fmt.Errorf("update game controller: %v", err)
		}

		ctx.String(http.StatusOK, "{\"success\": true}")
	case consts.ApplicationContentJSON:
		err := controller.Service.AttachSchedulesToGame(gameID)
		if err != nil {
			return fmt.Errorf("update game controller: %v", err)
		}
	default:
		ctx.String(http.StatusOK, httputils.GetBadRequestErrorResponseJSONStr())
	}

	return nil
}
