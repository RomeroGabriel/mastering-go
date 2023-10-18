package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error requesting:", err)
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		panic(err)
	}
	fmt.Println(string(res))
}
