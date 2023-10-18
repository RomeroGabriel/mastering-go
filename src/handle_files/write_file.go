package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("src/handle_files/example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	size, err := file.WriteString("\nWriting a string into the file\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Size written into file: ", size)
	_, err = file.Write([]byte("Using bytes to write"))
}
