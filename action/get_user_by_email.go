package action

import (
	"context"
	"log"

	"github.com/MuhAndriJP/user-service.git/entity"
	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/repo"
)

type GetUserByEmail struct {
	uRepo repo.User
}

func (u *GetUserByEmail) Handle(ctx context.Context, req *pb.GetUserByEmailRequest) (res *pb.GetUserResponse, err error) {
	res = &pb.GetUserResponse{}

	user, err := u.uRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	log.Println("user", user)
	if user != (entity.Users{}) {
		res = &pb.GetUserResponse{
			Name:  user.Name,
			Email: user.Email,
		}
	}
	log.Println("res", res)

	return
}

func NewGetUserByEmail() *GetUserByEmail {
	return &GetUserByEmail{
		uRepo: repo.NewUserRepo(),
	}
}
