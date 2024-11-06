package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	var err error
	go func() {
		// 启动grpc
		app := InitGrpcServer()
		err = app.server.Serve()
		if err != nil {
			panic(err)
		}
	}()

	// 初始化engin
	server := InitWebServer()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"hello": "world"})
		return
	})
	initPrometheus()
	// 指定端口
	err = server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

}
func initPrometheus() {
	go func() {
		// 专门给 prometheus 用的端口
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()
}
