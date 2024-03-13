package main

import (
	"fmt"
	"reflect"
)

func printSlice(s []int) {
	fmt.Println("printing s: ", s)
	fmt.Println("type of s: ", reflect.TypeOf(s))
}

func panicArray(data []int) {
	panicArray := [5]int(data)
	fmt.Println(panicArray)
}

func main() {
	arr := [3]int{1, 2, 3}
	printSlice(arr[:])

	slice := []int{1, 2, 3}
	var arrCopy [3]int
	copy(arrCopy[:], slice)
	fmt.Println("Slice to Array", arr)

	slice = []int{1, 2, 3, 4}
	array := [4]int(slice)
	arraySmall := [2]int(slice)
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(array)
	fmt.Println(arraySmall)
	panicArray(slice)
}
