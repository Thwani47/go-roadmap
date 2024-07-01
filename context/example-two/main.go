package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	fmt.Println("server: hello handler started")

	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintf(writer, "Hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server error:", err)
		internalErr := http.StatusInternalServerError
		http.Error(writer, err.Error(), internalErr)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
