package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameUpdateController struct {
	GameService               interfaces.IGameService
	AttachSourceToGameService interfaces.IAttachSourceToGameService
}

func (controller *GameUpdateController) UpdateGame(ctx echo.Context) error {
	gameID := ctx.Param("game-id")

	fmt.Println("GameUpdateController: GameID:", gameID)

	switch httputils.ParseContentType(ctx) {
	case consts.FormDataContentType:
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
			return fmt.Errorf("update game controller: %v", err)
		}

		game, err := controller.GameService.GetGame(gameID)
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

		err = controller.AttachSourceToGameService.AttachZipArchiveToGame(gameID, archives)
		if err != nil {
			return fmt.Errorf("update game controller: %v", err)
		}

		ctx.String(http.StatusOK, "{\"success\": true}")
	case consts.ApplicationContentJSON:
		err := controller.AttachSourceToGameService.AttachSchedulesToGame(gameID)
		if err != nil {
			return fmt.Errorf("update game controller: %v", err)
		}
	default:
		ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	return nil
}
