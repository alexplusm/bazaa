package controllers

import (
	"fmt"
	"net/http"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/errors"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/files"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
)

// TODO: move
const (
	// FormDataType const
	FormDataType = "multipart/form-data"
	// ApplicationJSON const
	ApplicationJSON = "application/json"
)

type schedule struct {
	ScheduleID string `json:"schedule_id" validate:"required"`
}

type updateGameWithSchedulesRequestBody struct {
	Schedules []schedule `json:"schedules" validate:"required,dive"`
}

// UpdateGame update game controller
func UpdateGame(p *pgxpool.Pool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("GAME ID", ctx.Param("game-id"))

		// TODO: check game existanse

		switch middlewares.ParseContentType(ctx) {
		case FormDataType:
			fmt.Println("FORM DATA TYPE")
			form, err := ctx.MultipartForm()
			if err != nil {
				// TODO: log error
				ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
				return fmt.Errorf("UpdateGame controller: %v", err)
			}

			/* Although the field is called "archives"
			* expected only one file - zip archive
			* The array is made for the future.
			 */
			archives := form.File["archives"]

			filenames, err := files.CopyFiles(archives, consts.MediaTempDir)
			if err != nil {
				fmt.Printf("Error while copieng: %+v\n", err) // TODO: log
			}

			fmt.Println("filenames:", filenames)

			// must return imageNames and category
			res, err := files.UnzipImages(filenames)
			if err != nil {
				fmt.Println("Error", err) // todo: process error!
			}

			fmt.Println("FILES", res, "| len:", len(res))
			fmt.Printf("\n\n\n")

			// todo: fill database use res

			removeArchives(filenames)
			ctx.String(http.StatusOK, "OKEY") // todo: {success: true}
		case ApplicationJSON:
			updateGameBody := new(updateGameWithSchedulesRequestBody)
			validate = validator.New()

			if err := ctx.Bind(updateGameBody); err != nil {
				return fmt.Errorf("UpdateGame controller: %v", err)
			}
			if err := validate.Struct(updateGameBody); err != nil {
				return fmt.Errorf("UpdateGame controller: %+v", err)
			}

			// todo: fill database with schedules
			fmt.Printf("UpdateGame! %+v:", updateGameBody)

		default:
			ctx.String(http.StatusOK, errors.GetBadRequestErrorResponseJSONStr())
		}
		return nil
	}
}

// fill db
// // LoadFilesToDBWrapper load to DB
// func LoadFilesToDBWrapper(p *pgxpool.Pool) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		r := files.FindAllImages()

// 		for _, ff := range r {
// 			img := dao.ImageDao{URL: ff, Category: "1"}
// 			dao.InsertImage(p, img)
// 		}
// 		return c.String(http.StatusOK, "loaded")
// 	}
// }

func removeArchives(filenames []string) {
	for _, fn := range filenames {
		if err := files.RemoveFile(consts.MediaTempDir, fn); err != nil {
			fmt.Println(err) // todo: log error
		}
	}
}
