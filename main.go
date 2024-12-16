package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/VashUber/go-grpc-http-service/internal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := internal.RunGRPC(ctx); err != nil {
			log.Fatalf("failed to run grpc: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := internal.RunHTTP(ctx); err != nil {
			log.Fatalf("failed to run http: %v\n", err)
		}
	}()

	wg.Wait()
}
