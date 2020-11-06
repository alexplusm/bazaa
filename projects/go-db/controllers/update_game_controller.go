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

		archives := form.File["archives"]

		err = controller.Service.AttachZipArchiveToGame(gameID, archives)
		if err != nil {
			return fmt.Errorf("update game controller: %v", err)
		}
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

/////////-----------
//func UpdateGame(p *pgxpool.Pool) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		switch httputils.ParseContentType(ctx) {
//		case consts.FormDataContentType:
//			fmt.Println("FORM DATA TYPE")
//			form, err := ctx.MultipartForm()
//			if err != nil {
//				// TODO:log // TODO:error
//				ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
//				return fmt.Errorf("UpdateGame controller: %v", err)
//			}
//
//			/* INFO: Although the field is called "archives" expected
//			 *only one file - zip archive. The array is made for the future.
//			 */
//			archives := form.File["archives"]
//
//			filenames, err := fileutils.CopyFiles(archives, consts.MediaTempDir)
//			if err != nil {
//				fmt.Printf("Error while copieng: %+v\n", err) // TODO:log
//			}
//
//			fmt.Println("filenames:", filenames)
//
//			// must return imageNames and category
//			res, err := fileutils.UnzipImages(filenames)
//			if err != nil {
//				fmt.Println("Error", err) // TODO:error
//			}
//
//			fmt.Println("FILES", res, "| len:", len(res))
//			fmt.Printf("\n\n\n")
//
//			// todo: fill database use res
//
//			removeArchives(filenames)
//			ctx.String(http.StatusOK, "OKEY") // todo: {success: true}
//		case consts.ApplicationContentJSON:
//			updateGameBody := new(updateGameWithSchedulesRequestBody)
//			validate = validator.New()
//
//			if err := ctx.Bind(updateGameBody); err != nil {
//				return fmt.Errorf("UpdateGame controller: %v", err)
//			}
//			if err := validate.Struct(updateGameBody); err != nil {
//				return fmt.Errorf("UpdateGame controller: %+v", err)
//			}
//
//			// todo: fill database with schedules
//			fmt.Printf("UpdateGame! %+v:", updateGameBody)
//
//		default:
//			ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
//		}
//		return nil
//	}
//}
//
