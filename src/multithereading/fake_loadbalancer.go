package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worked %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	workersNum := 1000

	for i := 0; i < workersNum; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 2000; i++ {
		ch <- i
	}
}
