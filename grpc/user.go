package grpc

import (
	userv1 "gin-micro-example/api/proto/gen/user/v1"
	"gin-micro-example/internal/domain"
	"gin-micro-example/internal/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type UserServiceServe struct {
	svc service.UserService
	userv1.UnimplementedUserServiceServer
}

func NewUserServiceServer(svc service.UserService) *UserServiceServe {
	return &UserServiceServe{svc: svc}
}

func (i *UserServiceServe) Register(s *grpc.Server) {
	userv1.RegisterUserServiceServer(s, i)
}

func (u *UserServiceServe) Get(ctx context.Context, request *userv1.UserRequest) (*userv1.UserResponse, error) {
	user, err := u.svc.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return u.toDTO(user), nil
}

func (u *UserServiceServe) toDTO(user domain.User) *userv1.UserResponse {
	return &userv1.UserResponse{
		Id:       user.Id,
		UserName: user.UserName,
		Age:      user.Age,
		IsActive: user.IsActive,
	}
}
