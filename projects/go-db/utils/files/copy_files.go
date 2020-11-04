package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// CopyFiles copy files to path
func CopyFiles(files []*multipart.FileHeader, copyPath string) ([]string, error) {
	filenames := make([]string, 0)
	for _, file := range files {
		fmt.Println("file:", file.Filename, file.Size)

		// Source
		src, err := file.Open()
		if err != nil {
			return filenames, err
		}
		defer src.Close()

		filename := filepath.Base(file.Filename)
		fp := filepath.Join(copyPath, filename) // TODO: validate "copyPath" ?

		// Destination
		dst, err := os.Create(fp)
		if err != nil {
			return filenames, err
		}
		defer dst.Close()

		// Copy
		if _, err := io.Copy(dst, src); err != nil {
			return filenames, err
		}
		filenames = append(filenames, filename)
	}
	return filenames, nil
}
