package internal

import (
	"context"
	"fmt"
	"net"

	foov1 "github.com/VashUber/go-grpc-http-service/gen/foo/v1"
	"google.golang.org/grpc"
)

func RunGRPC(ctx context.Context) error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}

	errCh := make(chan error)

	s := grpc.NewServer()
	foo := &FooService{}

	foov1.RegisterFooServiceServer(s, foo)

	go func() {
		if err := s.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	defer func() {
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Graceful shutdown grpc server...")
		s.GracefulStop()
	case err := <-errCh:
		return err
	}

	return nil
}
