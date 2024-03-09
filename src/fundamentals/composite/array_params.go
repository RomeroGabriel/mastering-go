package main

import "fmt"

func arrayParam(data [4]int) {
	fmt.Printf("Memory allocated for data: %p\n", &data)
	data[0] = 90
}

func main() {
	myArray := [4]int{1, 2, 3, 4}
	fmt.Println("Inital myArray: ", myArray)
	fmt.Printf("Memory allocated for myArray: %p\n", &myArray)
	arrayParam(myArray)
	fmt.Println("After call func Slice: ", myArray)
}
