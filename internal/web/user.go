package web

import (
	userv1 "gin-micro-example/api/proto/gen/user/v1"
	"gin-micro-example/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	svc       service.UserService
	remoteSvc userv1.UserServiceClient
}

func NewUserHandler(svc service.UserService, remoteSvc userv1.UserServiceClient) *UserHandler {
	return &UserHandler{
		svc:       svc,
		remoteSvc: remoteSvc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/hello", h.Hello)
	ug.GET("/:id", h.Get)
	ug.GET("/rpc/:id", h.GetByRpc)
}

func (h *UserHandler) Hello(ctx *gin.Context) {
	// 初始化随机种子
	// 将浮点数转换为time.Duration类型，并乘以time.Second以转换为纳秒
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, map[string]string{"hello": "user_world"})
	return
}

func (h *UserHandler) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	// 先忽略掉了错误
	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 先忽略掉了错误
	info, _ := h.svc.Get(ctx, id)
	ctx.JSON(http.StatusOK, info)
	return
}

func (h *UserHandler) GetByRpc(ctx *gin.Context) {
	idStr := ctx.Param("id")
	// 先忽略掉了错误
	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 先忽略掉了错误
	info, _ := h.remoteSvc.Get(ctx, &userv1.UserRequest{Id: id})
	ctx.JSON(http.StatusOK, info)
	return
}
