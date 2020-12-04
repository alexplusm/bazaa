package services

import (
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

type ImageService struct {
}

func (service *ImageService) BuildImageURL(imageName string) (string, error) {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT_OUTER")
	hostPort := host + ":" + port

	u, err := url.Parse(hostPort)
	if err != nil {
		return "", fmt.Errorf("build image url: %v", err)
	}

	u.Path = path.Join(consts.MediaUrlPart, imageName)

	return u.String(), nil
}
