package gogreenlight

import (
	"errors"
	"fmt"
)

type SemaphoreCollection map[string]*namedSemaphore

var Semaphores SemaphoreCollection

func init() {
	Semaphores = SemaphoreCollection{}
}

func (sc *SemaphoreCollection) Add(semaphore *namedSemaphore) {
	Semaphores[semaphore.name] = semaphore
}

func (sc *SemaphoreCollection) Get(name string) (*namedSemaphore, error) {
	semaphore, ok := Semaphores[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Named samaphore %s not found in collection", name))
	}
	return semaphore, nil
}
