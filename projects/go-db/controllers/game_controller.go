package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/logutils"
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
	if err != nil {
		log.Error("game details: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildInternalServerErrorResponse(),
		)
	}

	sourcesDTO := make([]dto.SourceDTO, 0, len(sources))
	for _, source := range sources {
		sourcesDTO = append(sourcesDTO, source.ToDTO())
	}

	body := game.ToDetailsDTO(sourcesDTO)

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(body))
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

func (controller *GameController) AttachArchives(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

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

	err = controller.AttachSourceToGameService.AttachArchives(gameID, archives)
	if err != nil {
		log.Error("game update controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponse(),
		)
	}

	return ctx.JSON(http.StatusOK, httputils.BuildSuccessWithoutBodyResponse())
}

func (controller *GameController) AttachSchedules(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	game, err := controller.GameService.Retrieve(gameID)
	if err != nil {
		log.Error("game update controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildErrorResponse(http.StatusOK, "game not found"),
		)
		// TODO: возвращать ошибку с payload и в httpErrorHandler обрабатывать и отправлять запрос!
	}
	if !game.NotStarted() {
		log.Info("game update controller: ", "game started: ", gameID)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildErrorResponse(http.StatusOK, "game started"),
		)
	}

	err = controller.AttachSourceToGameService.AttachSchedules(gameID)
	if err != nil {
		log.Error("game update controller: ", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponse(),
		)
	}

	return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
}

func (controller *GameController) AttachGameResults(ctx echo.Context) error {
	gameID := ctx.Param(consts.GameIDUrlParam)

	requestBody := dto.AttachGameResultsRequestBody{}
	if err := ctx.Bind(&requestBody); err != nil {
		log.Error(logutils.GetStructName(controller), "AttachGameResults:", err)
		return ctx.JSON(
			http.StatusOK, httputils.BuildBadRequestErrorResponse(),
		)
	}

	attachGameParams := bo.AttachGameParams{}
	attachGameParams.FromDTO(requestBody)

	game, err := controller.GameService.Retrieve(gameID)
	if err != nil {
		log.Error(logutils.GetStructName(controller), "AttachGameResults:", err)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildErrorResponse(http.StatusOK, "game not found"),
		)
	}

	if !game.NotStarted() {
		log.Error(logutils.GetStructName(controller), "AttachGameResults: game started", game)
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildErrorResponse(http.StatusOK, "game started"),
		)
	}

	exist, err := controller.SourceService.GameHasSomeSourceGameId(game.GameID, attachGameParams.SourceGameID)
	if err != nil {
		log.Error(logutils.GetStructName(controller), "AttachGameResults:", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}
	if exist {
		return ctx.JSON(
			http.StatusOK,
			httputils.BuildBadRequestErrorResponseWithMgs("game has some source game results"), // TODO: message
		)
	}

	err = controller.AttachSourceToGameService.AttachGameResults(game.GameID, attachGameParams)
	if err != nil {
		log.Error(logutils.GetStructName(controller), "AttachGameResults:", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	return nil
}

// TODO: !!!
// на форме !!!
// START_DATE, END_DATE, SCHEDULES_ID
