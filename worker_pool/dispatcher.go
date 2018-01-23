package worker_pool

import "go_rtb/internal/worker_pool/job"

var RtbApiCallJobQueue chan *job.DSPCallJob

type Dispatcher struct {
	MaxWorker  int
	WorkerPool chan chan *job.DSPCallJob
}

func NewDispatcher(maxWorker int) *Dispatcher {
	pool := make(chan chan *job.DSPCallJob, maxWorker)
	return &Dispatcher{
		WorkerPool: pool,
		MaxWorker:  maxWorker,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case j := <-RtbApiCallJobQueue:
			go func(j *job.DSPCallJob) {
				jobChannel := <-d.WorkerPool
				jobChannel <- j
			}(j)
		}
	}
}
