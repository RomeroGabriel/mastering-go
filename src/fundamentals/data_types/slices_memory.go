package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60}

	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	noneValue := mySlice[:0]
	fmt.Printf("noneValue: len=%d cap=%d %v\n", len(noneValue), cap(noneValue), noneValue)
	firstFour := mySlice[:4]
	fmt.Printf("firstFour: len=%d cap=%d %v\n", len(firstFour), cap(firstFour), firstFour)
	fmt.Println()
	mySlice = append(mySlice, 70)
	fmt.Println("mySlice with 70: ", mySlice)
	fmt.Printf("mySlice: len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("firstFour: len=%d cap=%d %v\n", len(firstFour), cap(firstFour), firstFour)
	fmt.Printf("noneValue: len=%d cap=%d %v\n", len(noneValue), cap(noneValue), noneValue)
}
