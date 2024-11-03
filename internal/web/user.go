package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/hello", h.Hello)
}

func (h *UserHandler) Hello(ctx *gin.Context) {
	// 初始化随机种子

	// 将浮点数转换为time.Duration类型，并乘以time.Second以转换为纳秒
	time.Sleep(time.Second)

	ctx.JSON(http.StatusOK, map[string]string{"hello": "user_world"})
	return
}
