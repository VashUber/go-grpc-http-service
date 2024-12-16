package internal

import (
	"context"
	"net/http"

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

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		return err
	}

	return nil
}
