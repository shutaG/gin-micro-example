package service

import (
	"gin-micro-example/internal/domain"
	"golang.org/x/net/context"
)

type UserService interface {
	Get(ctx context.Context, id int64) (domain.User, error)
}

type userService struct {
	a string
}

func NewUserService() UserService {
	return &userService{a: "1"}
}

func (u userService) Get(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{
		Id:       id,
		UserName: "张三",
		Age:      18,
		IsActive: true,
	}, nil
}
