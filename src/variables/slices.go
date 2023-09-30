package main

import "fmt"

func main() {
	my_slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Full slice: ", my_slice)
	first_three_elem := my_slice[0:3]
	fmt.Println("First 3 elements: ", first_three_elem)

	fmt.Printf("my_slice	= len=%d cap=%d %v\n", len(my_slice), cap(my_slice), my_slice)
	none_value := my_slice[:0]
	fmt.Printf("none_value 	= len=%d cap=%d %v\n", len(none_value), cap(none_value), none_value)
	first_four := my_slice[:4]
	fmt.Printf("first_four 	= len=%d cap=%d %v\n", len(first_four), cap(first_four), first_four)
	skip_two_first := my_slice[2:]
	fmt.Printf("last_two	= len=%d cap=%d %v\n", len(skip_two_first), cap(skip_two_first), skip_two_first)

	my_slice = append(my_slice, 70)
	fmt.Println("my_slice with 70: ", my_slice)
	fmt.Printf("my_slice	= len=%d cap=%d %v\n", len(my_slice), cap(my_slice), my_slice)
}
