package action

import (
	"context"
	"log"

	"github.com/MuhAndriJP/user-service.git/entity"
	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/middleware"
	"github.com/MuhAndriJP/user-service.git/repo"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	uRepo repo.User
}

func (u *LoginUser) Handle(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	res = &pb.LoginUserResponse{}

	user, err := u.uRepo.GetUserByEmail(ctx, req.Email)
	if err != nil || user == (entity.Users{}) {
		log.Println("User Not Found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("Wrong Password")
		return
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return
	}

	user.Token = token

	err = u.uRepo.UpsertUser(ctx, user)
	if err != nil {
		return
	}

	res = &pb.LoginUserResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return
}

func NewLoginUser() *LoginUser {
	return &LoginUser{
		uRepo: repo.NewUserRepo(),
	}
}
