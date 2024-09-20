package functions

import (
	"time"
)

type Timer struct {
	startTime time.Time
}

func NewTimer() *Timer {
	return &Timer{
		startTime: time.Now(),
	}
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() time.Duration {
	return time.Since(t.startTime)
}

func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.startTime)
}

func (t *Timer) ElapsedSeconds() float64 {
	return time.Since(t.startTime).Seconds()
}
