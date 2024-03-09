package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray [10]int

	for i := 0; i <= 9; i++ {
		value := (i + 1) * 10
		myArray[i] = value
	}

	copyedArray := myArray[:]
	fmt.Println("To create a copy of the array, use full slice omit both the start and end indices: ", copyedArray)
	fmt.Println("Notice that the copyedArray is actually a slice variable:", reflect.TypeOf(copyedArray))
	fmt.Printf("Address of myArray: %p\n", &myArray)
	fmt.Printf("Address of copyedArray: %p\n", &copyedArray)
	copyedArray[9] = 10000
	fmt.Println("Check that changing the copyedArray changes also the myArray!!!!!!!", copyedArray, myArray)
	fmt.Println()

	startToEnd := myArray[:5]
	fmt.Println("Getting 5 first elements, start to end: ", startToEnd)

	fromToEnd := myArray[5:]
	fmt.Println("Getting from index 5 to end, from to end: ", fromToEnd)

	betweenIndexes := myArray[2:5]
	fmt.Println("Getting between 2 and 5, between indexes: ", betweenIndexes)

	fmt.Println("Go does not have built-in support for stepping through a slice like Python does. However, you can achieve this by using a loop or by manually selecting the desired indices.")

	var steppedSlice []int
	for i := 0; i < len(myArray); i += 2 {
		steppedSlice = append(steppedSlice, myArray[i])
	}
	fmt.Println("Stepped Slice: ", steppedSlice)
}
