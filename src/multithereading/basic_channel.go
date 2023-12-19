package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello Word!"
	}()

	msg := <-channel
	fmt.Println(msg)
}
