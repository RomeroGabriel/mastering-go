package main

import "fmt"

func main() {
	strExample := "pi_π!"
	for i, r := range strExample {
		fmt.Println(i, r, string(r))
	}
	fmt.Println()
	fmt.Println("Notice that the index 4 was skipped and 960 in index 3 is bigger than a byte.")
	fmt.Println("Reason: π is a special char")
}
