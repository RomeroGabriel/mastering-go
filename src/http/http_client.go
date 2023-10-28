package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 3, // Set a timeout for the request
	}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
