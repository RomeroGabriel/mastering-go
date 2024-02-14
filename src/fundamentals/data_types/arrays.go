package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[1] = 10
	myArray[2] = 20
	fmt.Println("Indexes not set get default value: ", myArray[0])
	fmt.Println("All values in the array: ", myArray)
	fmt.Println()
	for i, value := range myArray {
		fmt.Printf("Value stored in %d index is %d\n", i, value)
	}

	fmt.Println()
	myArray[1] = 600
	fmt.Println("It's possible to change index value in array: ", myArray)
	fmt.Println()

	initArray := [3]string{"apple", "banana", "cherry"}
	for i, value := range initArray {
		fmt.Printf("Value stored in %d index is %v\n", i, value)
	}

	firstValue := initArray[1:]
	fmt.Println(firstValue)
}
