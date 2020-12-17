package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ImageFilterService struct {
	ValidateFacesService interfaces.IValidateFacesService
	ImageService         interfaces.IImageService
}

func (service *ImageFilterService) Filter() {

}
