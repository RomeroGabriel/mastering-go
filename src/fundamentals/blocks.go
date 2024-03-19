package main

import "fmt"

func main() {
	originalVar := 10
	if originalVar > 5 {
		fmt.Println(originalVar)
		originalVar := 7
		fmt.Println(originalVar)
	}
	fmt.Println(originalVar)

	int := "Redefining int type"
	fmt.Println(int)
}
