package mysql

import (
	"context"

	"github.com/MuhAndriJP/user-service.git/entity"
)

type MySQL interface {
	Insert(context.Context, *entity.Users) error
}

type SQL struct {
}

func (m *SQL) Insert(ctx context.Context, user *entity.Users) (err error) {
	if err = DB.Create(&user).Error; err != nil {
		return
	}
	return
}

func NewSQL() MySQL {
	return &SQL{}
}
