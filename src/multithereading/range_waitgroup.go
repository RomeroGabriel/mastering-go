package main

import (
	"fmt"
	"sync"
)

func publish(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Publishing value :%d\n", i)
		wg.Add(1) // using wg.Add(10) on line 27, this need to be removed
		ch <- i
	}
	close(ch)
}

func subscribe(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Subscribe value: :%d\n", x)
		wg.Done()
	}
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	// wg.Add(10)
	go publish(ch, &wg)
	subscribe(ch, &wg) // or go subscribe(ch, &wg) using wg.Add(10) on line 27
	wg.Wait()
}
