package controllers

import (
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
	AttachSourceToGameService interfaces.IAttachSourceToGameService
	ExtSystemService          interfaces.IExtSystemService
	GameService               interfaces.IGameService
	SourceService             interfaces.ISourceService
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

	gameID, err := controller.GameService.Create(*game)
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

	game, err := controller.GameService.Retrieve(gameID)
	if err != nil {
		log.Error("game details: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildNotFoundRequestErrorResponse("game not found"),
		)
	}
	sources, err := controller.SourceService.ListByGame(gameID)

	// TODO: service!!!

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

func (controller *GameController) List(ctx echo.Context) error {
	extSystemID := ctx.QueryParam(consts.ExtSystemIDQPName)

	exist, err := controller.ExtSystemService.Exist(extSystemID)
	if err != nil {
		log.Error("game list controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if !exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("extSystem not found"),
		)
	}

	gamesBO, err := controller.GameService.List(extSystemID)
	if err != nil {
		log.Error("game list controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	gamesDTO := make([]dto.GameItemResponseBody, 0, len(gamesBO))
	for _, game := range gamesBO {
		gamesDTO = append(gamesDTO, game.ToListItemDTO())
	}

	resp := dto.GameListResponseBody{Games: gamesDTO}
	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}

func (controller *GameController) Update(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	switch httputils.ParseContentType(ctx) {
	case consts.FormDataContentType:
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
		}

		game, err := controller.GameService.Retrieve(gameID)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildErrorResponse(http.StatusOK, "game not found"),
			)
		}

		if !game.NotStarted() {
			log.Info("game update controller: ", "game started: ", gameID)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildErrorResponse(http.StatusOK, "game started"),
			)
		}

		archives := form.File["archives"]

		if len(archives) == 0 {
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildBadRequestErrorResponseWithMgs("archive required"),
			)
		}

		err = controller.AttachSourceToGameService.AttachZipArchiveToGame(gameID, archives)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildBadRequestErrorResponse(),
			)
		}

		// TODO: return gameID?
		return ctx.JSON(http.StatusOK, httputils.BuildSuccessWithoutBodyResponse())
	case consts.ApplicationContentJSON:
		err := controller.AttachSourceToGameService.AttachSchedulesToGame(gameID)
		if err != nil {
			log.Error("game update controller: ", err)
			return ctx.JSON(
				http.StatusOK,
				httputils.BuildBadRequestErrorResponse(),
			)
		}
	}

	return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
}
