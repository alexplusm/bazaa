package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alexplusm/bazaa/projects/go-db/src/configs"
)

// FindAllImages find all images names
func FindAllImages() []string {
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
		if fileInfo.IsDir() {
			strs = findAllImagesPath(filepath.Join(fpath, fileInfo.Name()), strs)
		} else {
			strs = append(strs, fileInfo.Name())
		}
	}

	return strs
}
