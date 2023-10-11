package main

import (
	"fmt"
)

func FindMax[T int | float64](data []T) T {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	ints := []int{3, 7, 1, 9, 4, 6}
	floats := []float64{3.14, 2.71, 1.618, 2.718}

	fmt.Println("FindMax integer:", FindMax(ints))
	fmt.Println("FindMax float:", FindMax(floats))
}
