package server

import (
	"sync"
	"time"
)

type Limiter struct {
	requests int64
	rate     int64
	mutex    sync.Mutex
}

func NewLimiter(rate int64) *Limiter {
	return &Limiter{ rate: rate }
}

func (l Limiter) Run() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		l.mutex.Lock()
		l.requests = 0
		l.mutex.Unlock()
	}
}

func (l Limiter) IsAllowed() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.requests <= l.rate
}

func (l Limiter) AddRequest() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.requests  += 1
}
