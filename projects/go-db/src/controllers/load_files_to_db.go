package controllers

import (
	"github.com/jackc/pgx/pgxpool"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Alexplusm/bazaa/projects/go-db/src/configs"
	"github.com/Alexplusm/bazaa/projects/go-db/src/models"
	"github.com/labstack/echo"
)

// todo: remove from controllers?
func kek() []string {
	dir, err := os.Open(configs.MediaRoot)
	if err != nil {
		fmt.Println(err)
	}
	defer dir.Close()

	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}

	r := make([]string, 0, 20)
	for _, fileInfo := range filesInfo {
		// check existence Task with timestamp == "fileInfo.Name()".........
		fmt.Println("filesInfo", fileInfo.Name())

		r = findAllImagesPath(filepath.Join(configs.MediaRoot, fileInfo.Name()), r)

		fmt.Println("Result", r, len(r))
	}

	return r
}

func findAllImagesPath(fpath string, strs []string) []string {
	dir, err := os.Open(fpath)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer dir.Close()

	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, fileInfo := range filesInfo {
		if isDSStoreFile(fileInfo.Name()) {
			 continue
		}
		if fileInfo.IsDir() {
			strs = findAllImagesPath(filepath.Join(fpath, fileInfo.Name()), strs)
		} else {
			strs = append(strs, fileInfo.Name())
		}
	}

	return strs
}

// LoadFilesToDBWrapper load to DB
func LoadFilesToDBWrapper(p *pgxpool.Pool) func(echo.Context) error {
	return func (c echo.Context) error {
		r := kek()
	
		for _, ff := range r {
			img := models.ImageDao{URL: ff, Category: "1" }
			models.InsertImage(p, img)
		}
		return c.String(http.StatusOK, "loaded")
	}
}

// LoadFilesToDB load to DB
// func LoadFilesToDB(c echo.Context) error {
// 	r := kek()

// 	for _, ff := range r {
// 		img := models.ImageDao{URL: ff, Category: "1" }
// 		models.InsertImage(, img)
// 	}
// 	return c.String(http.StatusOK, "loaded")
// }
