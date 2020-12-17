package fileutils

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateDirIfNotExists(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Error("create dir: ", err)
	}
	// TODO: return err -> main -> FATAL
}

func RemoveFile(fpath string) error {
	if err := os.Remove(fpath); err != nil {
		return fmt.Errorf("remove file: %v", err)
	}
	return nil
}

func GetExtension(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[len(filenameParts)-1]
}
