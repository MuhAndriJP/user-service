package server

import (
	"context"

	"github.com/MuhAndriJP/user-service.git/action"
	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
)

type Server struct {
}

func (srv *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (res *pb.NoResponse, err error) {
	res = &pb.NoResponse{}
	if err = action.NewRegisterUser().Handle(ctx, req); err != nil {
		return
	}
	return
}

func (srv *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	res = &pb.LoginUserResponse{}
	res, err = action.NewLoginUser().Handle(ctx, req)
	return
}

func NewServerUser() *Server {
	return &Server{}
}
