package action

import (
	"context"

	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/repo"
)

type RegisterUser struct {
	uRepo repo.User
}

func (u *RegisterUser) Handle(ctx context.Context, req *pb.RegisterUserRequest) (err error) {
	err = u.uRepo.CreateUser(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return
	}

	return
}

func NewRegisterUser() *RegisterUser {
	return &RegisterUser{
		uRepo: repo.NewUserRepo(),
	}
}
