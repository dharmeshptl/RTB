package worker_pool

import (
	"go_rtb/internal/worker_pool/job"
)

type Worker struct {
	ID         int
	WorkerPool chan chan *job.DSPCallJob
	JobChanel  chan *job.DSPCallJob
	quit       chan bool
}

func NewWorker(id int, workPool chan chan *job.DSPCallJob) Worker {
	return Worker{
		ID:         id,
		WorkerPool: workPool,
		JobChanel:  make(chan *job.DSPCallJob),
		quit:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w *Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChanel

			select {
			case j := <-w.JobChanel:
				j.Process()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
