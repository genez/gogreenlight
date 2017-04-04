package gogreenlight

import (
	"errors"
	"fmt"
	"sync"
)

type SemaphoreCollection map[string]*namedSemaphore

var mutex sync.RWMutex
var Semaphores SemaphoreCollection

func init() {
	Semaphores = SemaphoreCollection{}
	mutex = sync.RWMutex{}
}

func (sc *SemaphoreCollection) Add(semaphore *namedSemaphore) {
	mutex.Lock()
	defer mutex.Unlock()

	Semaphores[semaphore.name] = semaphore
}

func (sc *SemaphoreCollection) Remove(name string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(Semaphores, name)
}

func (sc *SemaphoreCollection) Get(name string) (*namedSemaphore, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	semaphore, ok := Semaphores[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Named samaphore %s not found in collection", name))
	}
	return semaphore, nil
}
