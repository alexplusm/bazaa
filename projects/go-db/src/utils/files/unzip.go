package files

import (
	"strings"
)

// service files with this prefix appear after unpacking
const serviceUnzipFilePrefix = "._"

var extentionsWhiteList = [3]string{".jpg", "png", "jpeg"}

// IsInvalidImageFileName check image filename
func IsInvalidImageFileName(name string) bool {
	return strings.HasPrefix(name, serviceUnzipFilePrefix) ||
		!hasAllowableImageExtension(name)
}

func hasAllowableImageExtension(name string) bool {
	var availableExtention = false

	for _, extension := range extentionsWhiteList {
		if getExtension(name) == extension {
			availableExtention = true
		}
	}
	return availableExtention
}
