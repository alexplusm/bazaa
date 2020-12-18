package fileutils

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateDirIfNotExists(path string) error {
	return os.MkdirAll(path, 0777)
}

func RemoveFiles(filePaths []string) {
	for _, filePath := range filePaths {
		err := RemoveFile(filePath)
		if err != nil {
			log.Error("remove files: ", err)
		}
	}
}

func RemoveFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("remove file: %v", err)
	}
	return nil
}

func GetFileNames(files []*multipart.FileHeader) []string {
	fileNames := make([]string, 0, len(files))

	for _, file := range files {
		fileNames = append(fileNames, file.Filename)
	}

	return fileNames
}

func GetExtension(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[len(filenameParts)-1]
}
