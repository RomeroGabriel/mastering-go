package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ContextHandler(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// Could be ctx := context.Background()
	select {
	case <-time.After(5 * time.Second):
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Context completed successfully."))
		return
	case <-ctx.Done():
		log.Println("Request cancelled by client")
		return
	}
}

func ContextCancelHandler(res http.ResponseWriter, req *http.Request) {
	// Could be ctx := context.Background()
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	queryParams := req.URL.Query()
	number := 0
	if queryParams.Has("number") {
		num, err := strconv.Atoi(queryParams.Get("number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("ERROR Parsing number parameter to int"))
			return
		}
		number = num
	}

	if number%2 == 0 {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ContextCancel completed successfully."))
		return
	} else {
		cancel()
	}

	select {
	case <-ctx.Done():
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("ContextCancel Task canceled or timed out. Number cannot be odd"))
		return
	}
}

func ContextDeadLineHandler(res http.ResponseWriter, req *http.Request) {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(req.Context(), d)
	defer cancel()
	queryParams := req.URL.Query()
	number := 0
	if queryParams.Has("number") {
		num, err := strconv.Atoi(queryParams.Get("number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("ERROR Parsing number parameter to int"))
			return
		}
		number = num
	}

	time.Sleep(time.Duration(number) * time.Second)
	select {
	case <-ctx.Done():
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("ContextDeadLineHandler Task canceled or timed out."))
	default:
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ContextDeadLineHandler completed successfully."))
	}
	return
}

func ContextTimeoutHandler(res http.ResponseWriter, req *http.Request) {
	// Could be ctx := context.Background()
	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()

	queryParams := req.URL.Query()
	number := 0
	if queryParams.Has("number") {
		num, err := strconv.Atoi(queryParams.Get("number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("ERROR Parsing number parameter to int"))
			return
		}
		number = num
	}

	select {
	case <-time.After(time.Duration(number) * time.Second):
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ContextTimeout completed successfully."))
		return
	case <-ctx.Done():
		res.WriteHeader(http.StatusRequestTimeout)
		res.Write([]byte("ContextTimeout Task canceled or timed out."))
		return
	}
}

func ContextValueHandler(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	number := 0
	if queryParams.Has("number") {
		num, err := strconv.Atoi(queryParams.Get("number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("ERROR Parsing number parameter to int"))
			return
		}
		number = num
	}
	parentCtx, cancel := context.WithCancel(req.Context())
	defer cancel()
	// Could be ctx := context.Background()
	ctx := context.WithValue(parentCtx, "number", number)
	fmt.Println("Number passed", ctx.Value("number"))

	if number%2 == 0 {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ContextValue completed successfully"))
		return
	} else {
		cancel()
	}

	select {
	case <-ctx.Done():
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("ContextValue canceled or timed out. Number cannot be odd"))
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-context", ContextHandler)
	mux.HandleFunc("/get-context-cancel", ContextCancelHandler)
	mux.HandleFunc("/get-context-deadline", ContextDeadLineHandler)
	mux.HandleFunc("/get-context-timeout", ContextTimeoutHandler)
	mux.HandleFunc("/get-context-value", ContextValueHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server is listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
