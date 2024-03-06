package main

import "fmt"

func intFloatConversion() {
	fmt.Println("Example: Converting Int and Float")
	x := 10
	y := 15.5
	var sum1 float64 = y + float64(x)
	var sum2 int = x + int(y)
	fmt.Println("Sum1: ", sum1)
	fmt.Println("Sum2: ", sum2)
	fmt.Println()
}

func intByteConversion() {
	fmt.Println("Example: Converting Int and Byte")
	x := 10
	var y byte = 100
	var sum3 int = x + int(y)
	var sum4 byte = y + byte(x)
	fmt.Println("Sum3: ", sum3)
	fmt.Println("Sum4: ", sum4)
	fmt.Println()
}

func stringIntBoolComparison() {
	var str string = ""
	// if str -> will not work
	// if str == true -> will not work
	if str == "" {
		fmt.Println("Right way to compare string")
	}

	// if intNum -> will not work
	// if intNum == true -> will not work
	var intNum int = 1
	if intNum == 1 {
		fmt.Println("Right way to compare int")
	}
}

func main() {
	intFloatConversion()
	intByteConversion()
	stringIntBoolComparison()
}
