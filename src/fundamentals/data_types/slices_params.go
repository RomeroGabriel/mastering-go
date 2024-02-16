package main

import "fmt"

func sliceParam(data []int) {
	fmt.Printf("Memory allocated for data: %p\n", &data)
	data[0] = 90
	data = append(data, 90, 80, 70)
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Println("Inital mySlice: ", mySlice)
	fmt.Printf("Memory allocated for mySlice: %p\n", &mySlice)
	sliceParam(mySlice)
	fmt.Println("After call func Slice: ", mySlice)
	fmt.Println("Notice that append elements didn't change mySlice, just change value in index")
}
