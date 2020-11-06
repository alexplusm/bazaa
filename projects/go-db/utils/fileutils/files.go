package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateDirIfNotExists create dir
// todo: FileUtils os.MkdirAll(rootFolder, os.ModePerm)
// todo: os.MakeDirAll ???
func CreateDirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

// RemoveFile remove file
func RemoveFile(dir, filename string) error {
	fp := filepath.Join(dir, filename)
	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("RemoveFile: %v", err)
	}
	return nil
}

func getExtension(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[len(filenameParts)-1]
}
