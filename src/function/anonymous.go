package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	result1 := add(2, 4)
	fmt.Println("Result1: ", result1)
	result1 = add(3, 4)
	fmt.Println("Result1: ", result1)
	result1 = add(4, 4)
	fmt.Println("Result1: ", result1)

	result2 := func(a, b int) int {
		return a + b
	}(10, 10)
	fmt.Println("Result2: ", result2)
}
