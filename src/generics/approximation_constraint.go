package main

import (
	"fmt"
)

type NumberInt int
type NumberFloat float64

type ValidNumbers interface {
	~int | ~float64
}

func FindMax[T ValidNumbers](data []T) T {
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
	ints := []NumberInt{3, 7, 1, 9, 4, 6}
	floats := []NumberFloat{3.14, 2.71, 1.618, 2.718}

	fmt.Println("FindMax integer:", FindMax(ints))
	fmt.Println("FindMax float:", FindMax(floats))
}
