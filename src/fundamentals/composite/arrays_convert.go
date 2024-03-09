package main

import (
	"fmt"
	"reflect"
)

func printSlice(s []int) {
	fmt.Println("printing s: ", s)
	fmt.Println("type of s: ", reflect.TypeOf(s))
}

func main() {
	arr := [3]int{1, 2, 3}
	printSlice(arr[:])

	slice := []int{1, 2, 3}
	var arrCopy [3]int
	copy(arrCopy[:], slice)
	fmt.Println("Slice to Array", arr)
}
