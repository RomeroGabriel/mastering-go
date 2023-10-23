package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/api.project/address"
)

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

func GetAddressHandler(res http.ResponseWriter, req *http.Request) {
	queryData := req.URL.Query()
	if !queryData.Has("zipcode") {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	zipCodeParam := req.URL.Query().Get("zipcode")
	address, error := address.GetAddress(zipCodeParam)
	if error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(address)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-address", GetAddressHandler)
	mux_wrapped := setHeaders(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux_wrapped,
	}

	fmt.Println("Server is listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
