package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Value in index 2: ", mySlice[2])
	fmt.Println("All values in the slice: ", mySlice)
	fmt.Println()
	for i, value := range mySlice {
		fmt.Printf("Value stored in %d index is %d\n", i, value)
	}
	fmt.Println()
	mySlice2 := mySlice
	mySlice2[1] = 600
	fmt.Println("It's possible to change index value in array in both arrays: ", mySlice2, mySlice)

	mySlice3 := make([]int, len(mySlice))
	copy(mySlice3, mySlice)
	mySlice3[0] = 999
	fmt.Println("See that mySlice didn't change mySlice: ", mySlice3, mySlice)
}
