package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
