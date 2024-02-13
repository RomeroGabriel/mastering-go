package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[1] = 10
	myArray[2] = 20
	fmt.Println(myArray[0])
	fmt.Println(myArray)
	for i, value := range myArray {
		fmt.Printf("Value stored in %d index is %d\n", i, value)
	}

	initArray := [3]string{"apple", "banana", "cherry"}
	for i, value := range initArray {
		fmt.Printf("Value stored in %d index is %v\n", i, value)
	}
}
