package main

import "fmt"

func sum_check_even(num1, num2 int) (int, bool, error) {
	result := num1 + num2
	is_even := result%2 == 0
	return result, is_even, nil
}

func main() {
	result, is_even, _ := sum_check_even(1, 2)

	fmt.Printf("result is: %d\n", result)
	if is_even {
		fmt.Printf("result is even!!!\n")
	}

	result, is_even, _ = sum_check_even(2, 2)
	fmt.Printf("result is: %d\n", result)
	if is_even {
		fmt.Printf("result is even!!!\n")
	}

}
