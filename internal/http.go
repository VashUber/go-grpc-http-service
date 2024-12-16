package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	foov1 "github.com/VashUber/go-grpc-http-service/gen/foo/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTP(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := foov1.RegisterFooServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		return err
	}

	srv := http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	errCh := make(chan error)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	defer func() {
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("graceful shutdown http server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
	case err := <-errCh:
		return err
	}

	return nil
}
