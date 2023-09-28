package main

const A = "I'm a constant"

var B bool

func main() {
	println("Starting basic variables")
	print("Print A const: ")
	println(A)
	print("Print B var: ")
	println(B)
	println("NOTICE: B was infer as false")
	println()
	println("Different ways to declare a variable: ")
	var full_declare string = "A"
	println("Full declaration: " + full_declare)
	short_hand := "B"
	println("Short hand: " + short_hand)
}
