package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/anonymous_func", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from anonymous_func!"))
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
