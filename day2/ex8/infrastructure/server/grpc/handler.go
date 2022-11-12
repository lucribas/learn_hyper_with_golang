package grpc

import (
	"context"
	"fmt"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) GetByID(ctx context.Context, in *GetByIDInput) (*User, error) {
	return &User{
		Name: fmt.Sprintf("Hello, %s!", in.Id),
	}, nil
}
