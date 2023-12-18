package main

import (
	"fmt"
	"sync"
	"time"
)

func print_nice(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d from %s\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(24)

	go print_nice("Corinthians", &waitGroup)
	go print_nice("Knicks", &waitGroup)

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("Just an annoying print")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
