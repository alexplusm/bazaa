package files

import (
	"os"
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

func getExtension(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[len(filenameParts)-1]
}
