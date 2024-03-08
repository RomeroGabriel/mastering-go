package main

import "fmt"

func main() {
	var r1 rune = 'A'
	r2 := 'A'
	if r1 == r2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	str := "Hello,  世界" // The string contains both ASCII and non-ASCII characters.
	runes := []rune(str)
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	fmt.Println(string(runes))
}
