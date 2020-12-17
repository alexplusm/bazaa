package services

import (
	"fmt"
	"image"
	"net/url"
	"os"
	"path"

	"gocv.io/x/gocv"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
)

const (
	imageCropCoefficient = 0.09
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

func (service *ImageService) CropImage(imagePath string) error {
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		return fmt.Errorf("image empty: %v", imagePath)
	}

	size := img.Size()
	height := size[0]
	width := size[1]

	x0 := 0
	y0 := int(float64(height) * imageCropCoefficient)

	x1 := width
	y1 := height - y0

	croppedMat := img.Region(image.Rect(x0, y0, x1, y1))
	resultMat := croppedMat.Clone()

	gocv.IMWrite(imagePath, resultMat)
	//gocv.IMWrite(imagePath, croppedMat) // TODO: !!!

	return nil
}
