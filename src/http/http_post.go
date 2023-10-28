package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	payload := []byte(`{"key1": "value1", "key2": "value2"}`)
	body := bytes.NewBuffer(payload)

	resp, err := http.Post("https://httpbin.org/post", "application/json", body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
