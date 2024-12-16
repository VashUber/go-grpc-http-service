package internal

import (
	"context"

	foov1 "github.com/VashUber/go-grpc-http-service/gen/foo/v1"
)

type FooService struct {
	foov1.UnimplementedFooServiceServer
}

func (s *FooService) Ping(ctx context.Context, req *foov1.Id) (*foov1.Id, error) {
	return &foov1.Id{
		Id: req.Id + 1,
	}, nil
}
