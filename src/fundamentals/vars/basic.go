package main

import "fmt"

const A = "I'm a constant"

var B bool

func main() {
	fmt.Println("Starting basic variables")
	fmt.Println("Print A const: ", A)
	fmt.Println("Print B var: ", B)
	fmt.Println("NOTICE: B was infer as false")
	println()
	fmt.Println("Different ways to declare a variable: ")
	var full_declare string = "A"
	fmt.Println("Full declaration: ", full_declare)
	short_hand := "B"
	fmt.Println("Short hand: " + short_hand)
}
