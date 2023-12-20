package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)
			ch <- 1
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case num := <-ch:
			fmt.Printf("Number received: %d\n", num)
		default:
			log.Println("Default")
		}
	}
}
