package ratelimit

import (
	"errors"
	"sync"
	"time"
)

// Limiter controls how often actions can be executed
type Limiter struct {
	mu          sync.Mutex
	lastAction  time.Time
	minInterval time.Duration
	maxActions  int
	actionCount int
}

func NewLimiter(minInterval time.Duration, maxActions int) *Limiter {
	return &Limiter{
		minInterval: minInterval,
		maxActions:  maxActions,
	}
}

func (l *Limiter) Allow() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.actionCount >= l.maxActions {
		return errors.New("rate limit exceeded: daily quota reached")
	}

	if !l.lastAction.IsZero() {
		elapsed := time.Since(l.lastAction)
		if elapsed < l.minInterval {
			time.Sleep(l.minInterval - elapsed)
		}
	}

	l.lastAction = time.Now()
	l.actionCount++
	return nil
}
