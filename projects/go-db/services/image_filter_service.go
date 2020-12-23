package services

import (
	"archive/zip"
	"fmt"
	"path"
	"sync"
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
		//fmt.Println("worker", id, "processing job", j)

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

	workerCount := len(files)

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

// ---

type Job struct {
	ID        int
	Payload   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobChannel chan Job
type JobQueue chan chan Job

type Worker struct {
	ID      int
	JobChan JobChannel
	Queue   JobQueue
	Quit    chan struct{}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.Queue <- w.JobChan
			select {
			case job := <-w.JobChan:
				fmt.Println("payload", job)
				s := ValidateFacesService{}
				res, err := s.Validate(job.Payload)
				fmt.Println("RESULT: ", res, err)
			case <-w.Quit:
				close(w.JobChan)
				return
			}
		}
	}()
}

type Dispatcher struct {
	Workers  []*Worker
	WorkChan JobChannel
	Queue    JobQueue
}

func NewDispatcher(count int) *Dispatcher {
	return &Dispatcher{
		Workers:  make([]*Worker, count),
		WorkChan: make(JobChannel),
		Queue:    make(JobQueue),
	}
}

func (d *Dispatcher) Start() *Dispatcher {
	l := len(d.Workers)
	for i := 1; i <= l; i++ {
		wrk := &Worker{
			ID:      i,
			JobChan: make(JobChannel),
			Queue:   d.Queue,
			Quit:    make(chan struct{}),
		}
		wrk.Start()
		d.Workers = append(d.Workers, wrk)
	}
	go d.process()
	return d
}

func (d *Dispatcher) process() {
	for {
		select {
		case job := <-d.WorkChan:
			jobChan := <-d.Queue
			jobChan <- job
		}
	}
}

func (d *Dispatcher) Submit(job Job) {
	d.WorkChan <- job
}
