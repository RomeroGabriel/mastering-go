package main

import "fmt"

type MyInt int

func main() {
	var anyVar any
	var mInt MyInt = 20
	anyVar = mInt
	asser_int := anyVar.(MyInt)
	fmt.Println(asser_int + 1)

	assert_wrong, ok := anyVar.(string)
	if !ok {
		fmt.Println("Error")
	} else {
		// It will never print because this just happens in runtime and not compile time
		// Runtime error: panic: interface conversion: interface {} is main.MyInt, not string
		fmt.Println("THIS WILL NEVER PRINT: ", assert_wrong)
	}
}
