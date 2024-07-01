package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Thwani47/go-roadmap/context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx, "handler started!")
	defer log.Println(ctx, "handler ended!")

	fmt.Println(ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
