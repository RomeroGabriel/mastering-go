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

	rawStr := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.Nunc ac urna ex. Mauris eu dolor nec orci aliquam consectetur vitae ut sapien.
	Quisque sed commodo dui. Cras pulvinar aliquet eleifend. Nulla dignissim arcu quis nunc facilisis dictum.
Curabitur consequat, purus vel sodales dictum, dui metus lacinia lacus, in congue est.`
	fmt.Println(rawStr)
}
