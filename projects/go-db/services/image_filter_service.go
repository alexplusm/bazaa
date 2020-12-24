package services

import (
	"archive/zip"
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
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

	fmt.Println("Start filters", n, "len before: ", len(files))

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

	wp := workerpool.New(500)
	var mx sync.Mutex

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		wp.Submit(func() {
			ok, err := service.ValidateFacesService.Validate(filePath)
			if err != nil {
				// TODO: log error
				fmt.Println("Error while face validate")
			}
			if ok {
				mx.Lock()
				filteredFiles = append(filteredFiles, file)
				mx.Unlock()
			} else {
				// TODO: log error
				fmt.Println("File not valid", filePath)
			}
		})
	}

	wp.StopWait()

	return filteredFiles
}

func (service *ImageFilterService) cropFilter(files []zip.File) []zip.File {
	filteredFiles := make([]zip.File, 0, len(files))

	wp := workerpool.New(500)
	var mx sync.Mutex

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		wp.Submit(func() {
			err := service.ImageService.Crop(filePath)
			if err != nil {
				// TODO: log
				log.Error("error", err)
			}

			mx.Lock()
			filteredFiles = append(filteredFiles, file)
			mx.Unlock()
		})
	}

	return filteredFiles
}
