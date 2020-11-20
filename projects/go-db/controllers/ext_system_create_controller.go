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

type ExtSystemCreateController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *ExtSystemCreateController) CreateExtSystem(ctx echo.Context) error {
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
