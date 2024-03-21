package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// mutexes and atomic values

// state -> set or update or delete

type State struct {
	// mu    sync.RWMutex
	count int32
}

func (s *State) setState(i int) {
	atomic.AddInt32(&s.count, int32(i))
	// s.mu.Lock()
	// s.count = i
	// s.mu.Unlock()
}

func less_17() {
	state := &State{}
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		// doing some work here
		go func(i int) {
			state.setState(i + 1)
			waitGroup.Done() // done working
		}(i)
	}
	waitGroup.Wait()
	fmt.Printf("%+v\n", state)

}
