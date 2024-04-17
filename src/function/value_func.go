package main

import "fmt"

func test1(a string) int {
	fmt.Println("TEST1 A PARAM: ", a)
	return 10
}

func test2(b string) int {
	fmt.Println("TEST2 b PARAM: ", b)
	return 200
}

func main() {
	var funcSignature func(string) int
	funcSignature = test1

	result := funcSignature("CALLING funcSignature")
	fmt.Println("RESULT: ", result)

	funcSignature = test2
	result = funcSignature("CALLING AGAIN funcSignature")
	fmt.Println("RESULT: ", result)
}
