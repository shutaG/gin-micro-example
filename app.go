package main

import (
	"gin-micro-example/pkg/grpcx"
	"github.com/gin-gonic/gin"
)

type AppWebServe struct {
	server *gin.Engine
}

type AppGrpcServe struct {
	server *grpcx.Server
}
