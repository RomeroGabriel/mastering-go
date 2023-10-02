package main

import "fmt"

func main() {
	print_double(6)
	sum_result := sum(7, 3)
	fmt.Printf("sum_result is: %d\n", sum_result)
}

func print_double(double int) {
	result := double * 2
	fmt.Printf("Double of %d: %d\n", double, result)
}

func sum(num1, num2 int) int {
	return num1 + num2
}
