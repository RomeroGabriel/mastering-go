package main

import "fmt"

func print_double(num int) {
	num = num * 2
	fmt.Printf("	Memory allocated for num (inside print_double): %p\n", &num)
	fmt.Println("	Double: ", num)
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	num1 := 5
	fmt.Printf("Memory allocated for num1: %p\n", &num1)
	print_double(num1)
	fmt.Println("After call print_double: ", num1)
}
