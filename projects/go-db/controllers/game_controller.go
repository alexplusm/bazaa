package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type GameController struct {
	GameService   interfaces.IGameService
	SourceService interfaces.ISourceService
}

func (controller *GameController) Create(ctx echo.Context) error {
	gameRaw := new(dto.CreateGameRequestBody)

	if err := ctx.Bind(gameRaw); err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK, httputils.BuildBadRequestErrorResponse(),
		)
	}

	game := new(bo.GameBO)
	if err := game.FromDTO(*gameRaw, validate); err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("validation"),
		)
	}

	gameID, err := controller.GameService.CreateGame(*game)
	if err != nil {
		log.Error("game create controller: ", err)
		return ctx.JSON(
			http.StatusOK, httputils.BuildBadRequestErrorResponse(),
		)
	}

	resp := httputils.BuildSuccessResponse(dto.CreateGameResponseBody{GameID: gameID})
	return ctx.JSON(http.StatusOK, resp)
}

func (controller *GameController) Details(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	game, err := controller.GameService.GetGame(gameID)
	if err != nil {
		log.Error("game details: ", err)
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
