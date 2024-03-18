package main

import "fmt"

func main() {
	intSet := map[int]bool{
		5:  true,
		10: true,
		2:  true,
		8:  true,
		7:  true,
		3:  true,
		9:  true,
		1:  true,
	}

	if intSet[10] {
		fmt.Println("10 is in the set")
	}

	if intSet[20] {
		fmt.Println("20 is in the set")
	}

	setStruct := map[int]struct{}{
		5:  {},
		10: {},
		2:  {},
		8:  {},
		7:  {},
		3:  {},
		9:  {},
		1:  {},
	}

	if _, ok := setStruct[8]; ok {
		fmt.Println("8 is in the set")
	}

	if _, ok := setStruct[20]; ok {
		fmt.Println("20 is in the set")
	}
}
