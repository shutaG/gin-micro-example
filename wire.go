//go:build wireinject

package main

import (
	"gin-micro-example/internal/ioc"
	"gin-micro-example/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		web.NewUserHandler,
	)
	return gin.Default()
}
