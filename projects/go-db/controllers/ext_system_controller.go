package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ExtSystemController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *ExtSystemController) Create(ctx echo.Context) error {
	extSystemRaw := new(dto.CreateExtSystemRequestBody)

	if err := ctx.Bind(extSystemRaw); err != nil {
		log.Error("extSystem create controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	extSystem := new(bo.ExtSystemBO)
	if err := extSystem.FromDTO(*extSystemRaw, validate); err != nil {
		log.Error("extSystem create controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	extSystemID, err := controller.ExtSystemService.CreateExtSystem(*extSystem)
	if err != nil {
		log.Error("extSystem create controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildBadRequestErrorResponse())
	}

	return ctx.JSON(
		http.StatusOK,
		httputils.BuildSuccessResponse(dto.CreateExtSystemResponseBody{ID: extSystemID}),
	)
}

func (controller *ExtSystemController) List(ctx echo.Context) error {
	listBO, err := controller.ExtSystemService.ExtSystemList()
	if err != nil {
		log.Error("extSystem list controller: ", err)
		return ctx.JSON(http.StatusOK, httputils.BuildInternalServerErrorResponse())
	}

	listDTO := make([]dto.ExtSystemListItem, 0, len(listBO))
	for _, item := range listBO {
		listDTO = append(listDTO, item.ToDTO())
	}

	resp := dto.ExtSystemListResponseBody{ExtSystems: listDTO}
	return ctx.JSON(http.StatusOK, httputils.BuildSuccessResponse(resp))
}
