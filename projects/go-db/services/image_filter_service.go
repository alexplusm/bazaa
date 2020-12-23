package services

import (
	"archive/zip"
	"fmt"
	"path"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ImageFilterService struct {
	ValidateFacesService interfaces.IValidateFacesService
	ImageService         interfaces.IImageService
}

func (service *ImageFilterService) Filter(files []zip.File) []zip.File {
	filteredFiles := make([]zip.File, 0, len(files))

	n := time.Now().Unix()

	fmt.Println("!!!!!!!!", n, "len before: ", len(files))

	filteredFiles = service.faceFilter(files)

	faceTime := time.Now().Unix()
	fmt.Println("FACE TIME: ", faceTime, " | ", faceTime-n)

	filteredFiles = service.cropFilter(filteredFiles)

	cropTime := time.Now().Unix()
	fmt.Println("CROP TIME: ", cropTime, " | ", cropTime-faceTime)

	fmt.Println("TOTAL TIME: ", cropTime-n)

	fmt.Println("len after: ", len(filteredFiles))

	return filteredFiles
}

func (service *ImageFilterService) faceFilter(files []zip.File) []zip.File {
	filteredFiles := make([]zip.File, 0, len(files))

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		ok, err := service.ValidateFacesService.Validate(filePath)
		if err != nil {
			log.Error("error", err)
			continue
		}
		if !ok {
			log.Info("ne ok: todo")
			continue
		}

		filteredFiles = append(filteredFiles, file)
	}

	return filteredFiles
}

func (service *ImageFilterService) cropFilter(files []zip.File) []zip.File {
	filteredFiles := make([]zip.File, 0, len(files))

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		err := service.ImageService.Crop(filePath)
		if err != nil {
			log.Error("error", err)
			continue
		}

		filteredFiles = append(filteredFiles, file)
	}

	return filteredFiles
}
