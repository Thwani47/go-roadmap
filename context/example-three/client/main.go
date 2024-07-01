package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type key string

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, key("foo"), "bar")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9000", nil)

	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %d", resp.StatusCode)
	}

	io.Copy(os.Stdout, resp.Body)
}
