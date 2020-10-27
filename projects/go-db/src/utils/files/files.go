package files

import (
	"os"
)

// CreateDirIfNotExists create dir
// todo: os.MakeDirAll ???
func CreateDirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
