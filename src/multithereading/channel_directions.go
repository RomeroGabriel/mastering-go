package main

import "fmt"

func setData(ch chan<- string) {
	super_secret := "Corinthians"
	ch <- super_secret
}

func readData(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func main() {
	ch := make(chan string)
	go setData(ch)

	readData(ch)
}
