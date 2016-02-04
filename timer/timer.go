package timer

import (
	"time"
)

// Timer is the timer obeject
type Timer struct {
	ticker 	*time.Ticker
}

// newTimer return a new Timer object
func newTimer(min time.Duration) (*Timer, error) {
	return &Timer {
		ticker: time.NewTicker(min * time.Minute),
	}, nil
}

func (t *Timer) run() {
	go func(){
		for {
			select {
				case <- t.ticker.C:
					go func(){
						t.poll()
					}()
			}
		}
	}()
}

func (t *Timer) poll() {
	
}