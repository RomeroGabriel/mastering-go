package main

import "fmt"

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello")

	defer fmt.Println("Hi, testttting")
}
