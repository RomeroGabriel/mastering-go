package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func print_nice(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d from %s\n", i, name)
		time.Sleep(1 * time.Second)
		defer wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(14)

	go print_nice("Corinthians", &waitGroup)

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("Just an annoying print")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	done := make(chan bool)
	go func() {
		waitGroup.Wait()
		// Close the Channel
		close(done)
	}()
	select {
	case <-done:
		log.Println("All done")
	case <-time.After(9 * time.Second):
		log.Println("Hit timeout")
	}
}
