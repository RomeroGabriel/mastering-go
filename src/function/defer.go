package main

import "fmt"

func main() {
	data := 10
	defer func(v int) {
		fmt.Println("1º: ", v)
	}(data)
	data = 20
	defer func(v int) {
		fmt.Println("2º: ", v)
	}(data)
	data = 30
	fmt.Println("Exiting: ", data)
}
