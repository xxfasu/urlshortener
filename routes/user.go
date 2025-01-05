package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxfasu/urlshortener/internal/handler/v1/user_handler"
)

func userRouter(router *gin.Engine, handler *user_handler.Handler) {
	// 用户认证相关路由组 /api/auth
	u := router.Group("/api/auth")

	// 用户认证相关端点
	u.POST("/login", handler.Login)                  // 用户登录
	u.POST("/register", handler.Register)            // 用户注册
	u.POST("/forget", handler.ForgetPassword)        // 忘记密码
	u.GET("/register/:email", handler.SendEmailCode) // 发送注册验证码

}
