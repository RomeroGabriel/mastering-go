package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("src/handle_files/example.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
