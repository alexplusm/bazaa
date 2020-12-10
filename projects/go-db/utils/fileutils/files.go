package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateDirIfNotExists(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Error("create dir: ", err)
	}
	// TODO: return err -> main -> FATAL
}

func RemoveFile(dir, filename string) error {
	fp := filepath.Join(dir, filename)
	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("remove file: %v", err)
	}
	return nil
}

func GetExtension(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[len(filenameParts)-1]
}
