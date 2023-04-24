package server

import (
	"context"

	"github.com/MuhAndriJP/user-service.git/action"
	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
)

// Server is a GRPC Server for payment snap service
type Server struct {
}

// AuthAccessPublic handler
func (srv *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.NoResponse, error) {
	action.NewRegisterUser().Handle(ctx, req)
	return &pb.NoResponse{}, nil
}

// NewServer create new Server
func NewServerUser() *Server {
	return &Server{}
}
