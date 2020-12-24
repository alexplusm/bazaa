package workerpool

// TODO

// ---
//
//type Job struct {
//	ID        int
//	Payload   string
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}
//
//type JobChannel chan Job
//type JobQueue chan chan Job
//
//type Worker struct {
//	ID      int
//	JobChan JobChannel
//	Queue   JobQueue
//	Quit    chan struct{}
//}
//
//func (w *Worker) Start() {
//	go func() {
//		for {
//			w.Queue <- w.JobChan
//			select {
//			case job := <-w.JobChan:
//				fmt.Println("payload", job)
//				s := ValidateFacesService{}
//				res, err := s.Validate(job.Payload)
//				fmt.Println("RESULT: ", res, err)
//			case <-w.Quit:
//				close(w.JobChan)
//				return
//			}
//		}
//	}()
//}
//
//type Dispatcher struct {
//	Workers  []*Worker
//	WorkChan JobChannel
//	Queue    JobQueue
//}
//
//func NewDispatcher(count int) *Dispatcher {
//	return &Dispatcher{
//		Workers:  make([]*Worker, count),
//		WorkChan: make(JobChannel),
//		Queue:    make(JobQueue),
//	}
//}
//
//func (d *Dispatcher) Start() *Dispatcher {
//	l := len(d.Workers)
//	for i := 1; i <= l; i++ {
//		wrk := &Worker{
//			ID:      i,
//			JobChan: make(JobChannel),
//			Queue:   d.Queue,
//			Quit:    make(chan struct{}),
//		}
//		wrk.Start()
//		d.Workers = append(d.Workers, wrk)
//	}
//	go d.process()
//	return d
//}
//
//func (d *Dispatcher) process() {
//	for {
//		select {
//		case job := <-d.WorkChan:
//			jobChan := <-d.Queue
//			jobChan <- job
//		}
//	}
//}
//
//func (d *Dispatcher) Submit(job Job) {
//	d.WorkChan <- job
//}
