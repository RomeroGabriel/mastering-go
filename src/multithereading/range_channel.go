package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping 1"
		messages <- "ping 2"
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}
