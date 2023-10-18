package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("src/handle_files/example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File deleted successfully.")
}
