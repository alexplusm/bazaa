package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameInfoController struct {
	GameService   interfaces.IGameService
	SourceService interfaces.ISourceService
}

func (controller *GameInfoController) GetGameInfo(ctx echo.Context) error {
	gameID := ctx.Param("game-id")

	game, err := controller.GameService.GetGame(gameID)
	if err != nil {
		// TODO: log
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}
	sources, err := controller.SourceService.GetSourcesByGame(gameID)

	fmt.Println("Game: ", game)

	resp := dto.GameInfoResponseBody{}
	resp.StartDate = strconv.FormatInt(game.StartDate.Unix(), 10)
	resp.FinishDate = strconv.FormatInt(game.EndDate.Unix(), 10)

	question := dto.QuestionDTO{}
	question.AnswerType = game.AnswerType
	question.Text = game.Question

	optsList := make([]dto.OptionDTO, 0, 10)
	opts := strings.Split(game.Options, ",")
	for i, o := range opts {
		optsList = append(optsList, dto.OptionDTO{Option: i, Text: o})
	}
	question.Options = optsList

	resp.Question = question

	sourcesDTO := make([]dto.SourceDTO, 0, len(sources))
	for _, s := range sources {
		// TODO:!!!! s.Type -> "file"
		sourcesDTO = append(sourcesDTO, dto.SourceDTO{Type: "file", SourceID: s.SourceID})
	}
	resp.Sources = sourcesDTO

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
