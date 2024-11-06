//go:build wireinject

package main

import (
	"gin-micro-example/grpc"
	"gin-micro-example/internal/ioc"
	"gin-micro-example/internal/service"
	"gin-micro-example/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		ioc.InitUserClient,
		web.NewUserHandler,
		service.NewUserService,
	)
	return gin.Default()
}
func InitGrpcServer() *AppGrpcServe {
	wire.Build(
		service.NewUserService,
		grpc.NewUserServiceServer,
		ioc.NewGrpcServer,
		wire.Struct(new(AppGrpcServe), "*"),
	)
	return &AppGrpcServe{}
}
