package ioc

import (
	grpc2 "gin-micro-example/grpc"
	"gin-micro-example/pkg/grpcx"
	"google.golang.org/grpc"
)

func NewGrpcServer(userSvc *grpc2.UserServiceServe) *grpcx.Server {
	s := grpc.NewServer()
	userSvc.Register(s)
	return &grpcx.Server{
		Server: s,
		Addr:   ":8090",
	}
}
