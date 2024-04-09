package main

import "fmt"

func sum_check_even(num1, num2 int) (sum_result int, is_even bool, _ error) {
	sum_result = num1 + num2
	is_even = sum_result%2 == 0
	return sum_result, is_even, nil
}

func sum2(num1, num2 int) (sum_result int, is_even bool, _ error) {
	sum_result = num1 + num2
	is_even = sum_result%2 == 0
	return 100, false, nil
}

// NEVER USER BLANK RETURN
func sum_blank(num1, num2 int) (sum_result int, is_even bool) {
	sum_result = num1 + num2
	is_even = sum_result%2 == 0
	return
}

func main() {
	var num = 3
	result, is_even, _ := sum_check_even(num, num)
	fmt.Println("result: ", result, " is_even: ", is_even)

	result, is_even, _ = sum2(num, num)
	fmt.Println("result: ", result, " is_even: ", is_even)

	result, is_even = sum_blank(num, num)
	fmt.Println("result: ", result, " is_even: ", is_even)
}
