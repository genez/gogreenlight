package gogreenlight

import (
	"context"
	"sync"
)

func NewNamedSemaphore(name string) *namedSemaphore {
	s := &namedSemaphore{
		ctx:  context.Background(),
		name: name,
		set:  false,
		c:    sync.NewCond(&sync.Mutex{}),
	}
	Semaphores.Add(s)
	return s
}

func NewNamedSemaphoreWithContext(name string, ctx context.Context) *namedSemaphore {
	s := &namedSemaphore{
		ctx:  ctx,
		name: name,
		set:  false,
		c:    sync.NewCond(&sync.Mutex{}),
	}
	Semaphores.Add(s)
	return s
}
