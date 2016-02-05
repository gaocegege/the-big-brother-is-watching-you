package timer

import (
	"log"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/worker"
)

// Timer is the timer obeject
type Timer struct {
	ticker *time.Ticker
	worker *worker.Worker
}

// NewTimer return a new Timer object
func NewTimer(min time.Duration, worker *worker.Worker) (*Timer, error) {
	return &Timer{
		ticker: time.NewTicker(min * time.Minute),
		worker: worker,
	}, nil
}

// Run the timer
func (t *Timer) Run() {
	log.Print("Timer is running now.")
	go func() {
		for {
			select {
			case <-t.ticker.C:
				go func() {
					log.Print("Worker is doing its jobs.")
					t.worker.Work()
				}()
			}
		}
	}()
}
