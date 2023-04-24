package repo

import (
	"context"

	"github.com/MuhAndriJP/user-service.git/entity"
	"github.com/MuhAndriJP/user-service.git/repo/mysql"
)

type User interface {
	CreateUser(context.Context, string, string, string) error
}

type UserRepo struct {
	dbSQL mysql.MySQL
}

func (r *UserRepo) CreateUser(ctx context.Context, name, email, password string) (err error) {
	if err = r.dbSQL.Insert(ctx, &entity.Users{
		Name:     name,
		Email:    email,
		Password: password,
	}); err != nil {
		return
	}

	return
}

func NewUserRepo() User {
	return &UserRepo{
		dbSQL: mysql.NewSQL(),
	}
}
