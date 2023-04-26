package repo

import (
	"context"

	"github.com/MuhAndriJP/user-service.git/entity"
	"github.com/MuhAndriJP/user-service.git/repo/mysql"
)

type User interface {
	CreateUser(context.Context, string, string, string) error
	GetUserByEmail(context.Context, string) (entity.Users, error)
	UpsertUser(context.Context, entity.Users) error
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

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (res entity.Users, err error) {
	res, err = r.dbSQL.GetUserByEmail(ctx, email)
	if err != nil {
		return
	}
	return
}

func (r *UserRepo) UpsertUser(ctx context.Context, user entity.Users) (err error) {
	if err = r.dbSQL.Upsert(ctx, &user); err != nil {
		return
	}
	return
}

func NewUserRepo() User {
	return &UserRepo{
		dbSQL: mysql.NewSQL(),
	}
}
