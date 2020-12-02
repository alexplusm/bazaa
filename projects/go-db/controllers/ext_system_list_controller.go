package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/httputils"
)

type ExtSystemListController struct {
	ExtSystemService interfaces.IExtSystemService
}

func (controller *ExtSystemListController) List(ctx echo.Context) error {
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
