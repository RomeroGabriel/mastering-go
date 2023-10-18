package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("src/handle_files/example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
