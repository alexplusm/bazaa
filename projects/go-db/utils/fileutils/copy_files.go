package fileutils

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

		filename := filepath.Base(file.Filename)
		fp := filepath.Join(copyPath, filename) // TODO: validate "copyPath" ?

		// Destination
		dst, err := os.Create(fp)
		if err != nil {
			return filenames, err
		}

		// Copy
		if _, err := io.Copy(dst, src); err != nil {
			return filenames, err
		}
		filenames = append(filenames, filename)

		if err = src.Close(); err != nil {
			return filenames, err
		}
		if err = dst.Close(); err != nil {
			return filenames, err
		}

	}
	return filenames, nil
}
