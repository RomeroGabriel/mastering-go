package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "message 1"
	channel <- "message 2"

	msg1 := <-channel
	msg2 := <-channel
	fmt.Println(msg1)
	fmt.Println(msg2)
}
