package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Basic for type 1: ", i)
	}
	fmt.Println("*********")

	j := 2
	for j < 5 {
		fmt.Println("Basic for type 2: ", j)
		j++
	}
	fmt.Println("*********")

	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println("arr for, index: ", index, "value ", value)
	}
	fmt.Println("*********")

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range myMap {
		fmt.Println("map for, key: ", key, "value ", value)
	}
}
