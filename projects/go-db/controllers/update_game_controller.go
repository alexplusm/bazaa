package controllers

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type UpdateGameController struct {
	Service interfaces.IUpdateGameService
}

func (controller *UpdateGameController) UpdateGame(ctx echo.Context) error {
	gameID := ctx.Param("game-id")

	//switch httputils.ParseContentType(ctx) {
	//
	//}

	fmt.Println("UpdateGameController: GameID:", gameID)

	// TODO: check game game existance in service

	err := controller.Service.AttachZipArchiveToGame(gameID)
	if err != nil {
		return fmt.Errorf("update game controller: %v", err)
	}

	return nil
}

/////////-----------

//type schedule struct {
//	ScheduleID string `json:"schedule_id" validate:"required"`
//}
//
//type updateGameWithSchedulesRequestBody struct {
//	Schedules []schedule `json:"schedules" validate:"required,dive"`
//}
//
//// UpdateGame update game controller
//func UpdateGame(p *pgxpool.Pool) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		fmt.Println("GAME ID", ctx.Param("game-id"))
//
//		// TODO: check game existanse
//
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
//func removeArchives(filenames []string) {
//	for _, fn := range filenames {
//		if err := fileutils.RemoveFile(consts.MediaTempDir, fn); err != nil {
//			fmt.Println(err) // TODO:log // TODO:error
//		}
//	}
//}
