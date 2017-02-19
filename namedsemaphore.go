package gogreenlight

import (
	"context"
	"sync"
	"time"
)

type namedSemaphore struct {
	c    *sync.Cond
	set  bool
	ctx  context.Context
	name string
}

func (s *namedSemaphore) Set() bool {
	s.c.L.Lock()
	defer s.c.L.Unlock()

	if !s.set {
		s.set = true
		s.c.Broadcast()
		return true
	} else {
		return false
	}
}

func (s *namedSemaphore) Unset() bool {
	s.c.L.Lock()
	defer s.c.L.Unlock()

	if s.set {
		s.set = false
		return true
	} else {
		return false
	}
}

func (s *namedSemaphore) Wait() bool {
	s.c.L.Lock()
	defer s.c.L.Unlock()
	select {
	case <-doWaitAsync(s.c):
		return true
	case <-s.ctx.Done():
		return false
	}
}

func (s *namedSemaphore) WaitWithTimeout(duration time.Duration) bool {
	s.c.L.Lock()
	defer s.c.L.Unlock()
	select {
	case <-doWaitAsync(s.c):
		return true
	case <-time.After(duration):
		return false
	case <-s.ctx.Done():
		return false
	}
}

func doWaitAsync(cond *sync.Cond) <-chan time.Time {
	ch := make(chan time.Time)
	go func(c *sync.Cond) {
		c.Wait()
		ch <- time.Now()
	}(cond)
	return ch
}
