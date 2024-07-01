package main

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

func enrichContext(ctx context.Context) context.Context {
	key := contextKey("request-id")
	return context.WithValue(ctx, key, "1234")
}

func doSomething(ctx context.Context) {
	key := contextKey("request-id")
	requestId := ctx.Value(key)
	fmt.Println(requestId)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timedout...")
			return
		default:
			fmt.Println("doing something...")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("Hello Context")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	fmt.Println(ctx.Err())

	doSomething(ctx)

	<-ctx.Done()
	fmt.Println("main timedout...")
	fmt.Println(ctx.Err())

	time.Sleep(2 * time.Second)
}
