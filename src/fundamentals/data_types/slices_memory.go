package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60}

	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	noneValue := mySlice[:0]
	fmt.Printf("noneValue: len=%d cap=%d %v\n", len(noneValue), cap(noneValue), noneValue)
	firstFour := mySlice[:4]
	fmt.Printf("firstFour: len=%d cap=%d %v\n", len(firstFour), cap(firstFour), firstFour)
	skipTwoFirst := mySlice[2:]
	fmt.Printf("last_two: len=%d cap=%d %v\n", len(skipTwoFirst), cap(skipTwoFirst), skipTwoFirst)

	mySlice = append(mySlice, 70)
	fmt.Println("mySlice with 70: ", mySlice)
	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
}
