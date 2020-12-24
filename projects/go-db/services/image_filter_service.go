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

	fmt.Println("!!!!!!!!", n, "len before: ", len(files))

	//filteredFiles = service.faceFilter(files)
	filteredFiles = service.faceFilterN(files)
	//filteredFiles = service.faceFilterO(files)

	faceTime := time.Now().Unix()
	fmt.Println("FACE TIME: ", faceTime, " | ", faceTime-n)

	filteredFiles = service.cropFilter(filteredFiles)

	cropTime := time.Now().Unix()
	fmt.Println("CROP TIME: ", cropTime, " | ", cropTime-faceTime)

	fmt.Println("TOTAL TIME: ", cropTime-n)

	fmt.Println("len after: ", len(filteredFiles))

	return filteredFiles
}

// --- old
func (service *ImageFilterService) faceFilterO(files []zip.File) []zip.File {
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

// ---

func worker(
	id int,
	wg *sync.WaitGroup,
	mx *sync.Mutex,
	jobs <-chan string,
	result map[string]bool,
	errors chan<- error,
) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)

		// TODO: !!!

		s := ValidateFacesService{}
		res, err := s.Validate(j)

		if err != nil {
			errors <- fmt.Errorf("error on job: %v\n", j)
		} else {
			mx.Lock()
			result[j] = res
			mx.Unlock()
		}

		wg.Done()
	}
}

func (service *ImageFilterService) faceFilter(files []zip.File) []zip.File {
	filteredFiles := make([]zip.File, 0, len(files))

	//  TODO: remove?
	filterMap := make(map[string]bool)

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		filterMap[filePath] = false
	}
	//  TODO: remove?

	var wg sync.WaitGroup
	var mx sync.Mutex

	jobs := make(chan string, len(files))
	errors := make(chan error, len(files))

	//workerCount := len(files)
	workerCount := 1000

	for w := 1; w <= workerCount; w++ {
		go worker(w, &wg, &mx, jobs, filterMap, errors)
	}

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())
		jobs <- filePath
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()

	for key, value := range filterMap {
		if value {
			for _, file := range files {
				filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())
				if filePath == key {
					filteredFiles = append(filteredFiles, file)
					continue
				}
			}
		}
	}

	//wp := workerpool.New(100)

	return filteredFiles
}

// ---

func (service *ImageFilterService) faceFilterN(files []zip.File) []zip.File {
	//filteredFiles := make([]zip.File, 0, len(files))
	//
	////  TODO: remove?
	//filterMap := make(map[string]bool)
	//
	//for _, file := range files {
	//	filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())
	//
	//	filterMap[filePath] = false
	//}
	////  TODO: remove?
	//
	//var wg sync.WaitGroup
	//var mx sync.Mutex
	//
	//jobs := make(chan string, len(files))
	//errors := make(chan error, len(files))
	//
	////workerCount := len(files)
	//workerCount := 1000

	//for w := 1; w <= workerCount; w++ {
	//	go worker(w, &wg, &mx, jobs, filterMap, errors)
	//}
	//
	//for _, file := range files {
	//	filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())
	//	jobs <- filePath
	//	wg.Add(1)
	//}
	//close(jobs)
	//
	//wg.Wait()

	//for key, value := range filterMap {
	//	if value {
	//		for _, file := range files {
	//			filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())
	//			if filePath == key {
	//				filteredFiles = append(filteredFiles, file)
	//				continue
	//			}
	//		}
	//	}
	//}

	filteredFiles := make([]zip.File, 0, len(files))

	wp := workerpool.New(100)

	for _, file := range files {
		filePath := path.Join(consts.MediaRoot, file.FileInfo().Name())

		wp.Submit(func() {
			ok, err := service.ValidateFacesService.Validate(filePath)
			if err != nil {
				// log error
				fmt.Println("Error while face validate")
			}
			if ok {
				filteredFiles = append(filteredFiles, file)
			} else {
				// log
				fmt.Println("File not valid", filePath)
			}
		})
	}

	wp.StopWait()

	return filteredFiles
}

// ---

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
