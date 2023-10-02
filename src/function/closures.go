package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	result1 := add(2, 4)
	fmt.Printf("result1 %d\n", result1)
	result1 = add(3, 4)
	fmt.Printf("result1 %d\n", result1)
	result1 = add(4, 4)
	fmt.Printf("result1 %d\n", result1)
}
