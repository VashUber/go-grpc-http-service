package internal

import (
	"context"
	"net"

	foov1 "github.com/VashUber/go-grpc-http-service/gen/foo/v1"
	"google.golang.org/grpc"
)

func RunGRPC(ctx context.Context) error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	foo := &FooService{}

	foov1.RegisterFooServiceServer(s, foo)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
