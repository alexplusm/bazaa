package services

import (
	"net"
	"net/url"
	"os"
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

type ImageService struct {
}

func (service *ImageService) BuildImageURL(imageName string) string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT_OUTER")
	hostPort := net.JoinHostPort(host, port)

	u, err := url.Parse(hostPort)
	if err != nil {
		log.Error("build image url: ", err)
	}

	u.Path = path.Join(hostPort, consts.MediaUrlPart, imageName)

	return u.String()
}
