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
	if !s.set {
		s.set = true
		s.c.Broadcast()
		return true
	} else {
		s.c.Broadcast()
		return false
	}
}

func (s *namedSemaphore) Unset() bool {
	if s.set {
		s.set = false
		return true
	} else {
		return false
	}
}

func (s *namedSemaphore) Wait() bool {
	select {
	case <-doWaitAsync(s):
		return true
	case <-s.ctx.Done():
		return false
	}
}

func (s *namedSemaphore) WaitWithContext(ctx context.Context) bool {
	select {
	case <-doWaitAsync(s):
		return true
	case <-ctx.Done():
		return false
	}
}

func (s *namedSemaphore) WaitWithTimeout(duration time.Duration) bool {
	select {
	case <-doWaitAsync(s):
		return true
	case <-time.After(duration):
		return false
	case <-s.ctx.Done():
		return false
	}
}

func doWaitAsync(s *namedSemaphore) <-chan time.Time {
	ch := make(chan time.Time)
	go func(ns *namedSemaphore) {
		ns.c.L.Lock()
		for !ns.set {
			ns.c.Wait()
		}
		ns.c.L.Unlock()
		ch <- time.Now()
	}(s)
	return ch
}
