package main

import "fmt"

func main() {
	str1 := "Corinthians"
	str2 := "Corinthians"
	fmt.Println("Memory allocated for str1: ", &str1)
	fmt.Println("Memory allocated for str2: ", &str2)
	fmt.Println("Even though both have the same value, the memory address is different.")
	fmt.Println()

	save := []byte("Salve o ")
	saveStr1 := append(save, str1...)
	fmt.Println("Using append bytes: ", string(saveStr1))

	strCopy := make([]byte, len(save))
	copy(strCopy, save)
	fmt.Println("Using copy bytes: ", string(strCopy))
}
