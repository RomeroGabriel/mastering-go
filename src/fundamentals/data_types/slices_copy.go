package main

import "fmt"

func main() {
	source := []int{1, 2, 3, 4}

	destination := make([]int, 2)
	num := copy(destination, source)
	fmt.Println("Copying first 2 elements", destination, num)

	copy(destination, source[2:])
	fmt.Println("Copying last 2 elements", destination, num)

	num = copy(source[:3], source[1:])
	fmt.Println("Overlaping elements in source: ", source, num)
	fmt.Println()
	fmt.Println("Using arrays: ")

	source = []int{1, 2, 3, 4}
	array := [4]int{5, 6, 7, 8}
	destination = make([]int, 2)
	copy(destination, array[:]) // create a slice from array
	fmt.Println("Copy from array", destination)
	copy(array[:], source)
	fmt.Println("Copying all elements from source to array: ", array)
}
