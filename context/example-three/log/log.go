package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type requestIDKey string

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey("requestId")).(int64)

	if !ok {
		log.Println("Could not find request iD in context")
		return
	}

	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := rand.Int63()

		ctx = context.WithValue(ctx, requestIDKey("requestId"), id)

		f(w, r.WithContext(ctx))
	}
}
