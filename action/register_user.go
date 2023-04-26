package action

import (
	"context"

	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/repo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	uRepo repo.User
}

func (u *RegisterUser) Handle(ctx context.Context, req *pb.RegisterUserRequest) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = u.uRepo.CreateUser(ctx, req.Name, req.Email, string(hashedPassword))
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
