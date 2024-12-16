package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/VashUber/go-grpc-http-service/internal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := internal.RunGRPC(ctx); err != nil {
			log.Fatalf("failed to run grpc: %v", err)
		}
	}()

	go func() {
		if err := internal.RunHTTP(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to run http: %v\n", err)
		}
	}()

	<-ctx.Done()
}
