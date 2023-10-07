package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Value of x: ", x)

	ptr := &x
	fmt.Println("Value that ptr point to: ", *ptr)
	fmt.Println("Memory address of ptr: ", ptr)
	*ptr = 100
	fmt.Println("Updated value of x: ", x)
}
