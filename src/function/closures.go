package main

import "fmt"

func closure1() {
	data := 10
	func() {
		fmt.Println("Inner Function -> closure1: ", data)
		data = 30
	}()
	fmt.Println("closure1 ENDING: ", data)
}

func closure2() {
	data := 10
	func() {
		fmt.Println("Inner Function -> closure2: ", data)
		data := 30
		fmt.Println("Inner Function SHADOW DATA -> closure2: ", data)
	}()
	fmt.Println("closure2 ENDING: ", data)
}

func main() {
	closure1()
	fmt.Println()
	closure2()
}
