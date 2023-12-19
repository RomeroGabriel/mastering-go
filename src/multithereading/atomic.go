package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func increment() {
	atomic.AddInt32(&counter, 1)
	wg.Done()
}

func main() {
	const numGoroutines = 1000
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println("Counter: ", counter)
}
