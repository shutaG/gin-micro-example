package ioc

import (
	"gin-micro-example/internal/web"
	proMid "gin-micro-example/pkg/prometheus"
	"github.com/gin-gonic/gin"
)

func InitWebServer(middlewares []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares() []gin.HandlerFunc {
	pm := &proMid.MiddlewareBuilder{
		Namespace: "ashutaG",
		Subsystem: "gin_micro_example",
		Name:      "gin_http",
		Help:      "统计HTTP接口数据",
	}

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			println("这是我的 Middleware")
		},
		pm.BuildResponseTime(),
		pm.BuildActiveRequest(),
	}
}
