package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	payload := []byte(`{"key1": "value1", "key2": "value2"}`)
	body := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", "https://httpbin.org/post", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
