package main

import "fmt"

func CheckNumbers[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func FindIndex[T comparable](data []T, value T) int {
	for i, data_value := range data {
		if data_value == value {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("CheckNumbers: ", CheckNumbers(10, 10.5))
	fmt.Println("CheckNumbers: ", CheckNumbers(10.5, 10.5))

	ints := []int{3, 7, 1, 9, 4, 6}
	fmt.Println("FindMax integer:", FindIndex(ints, 300))
	fmt.Println("FindMax integer:", FindIndex(ints, 9))
	fmt.Println("FindMax integer:", FindIndex(ints, 10))
}
