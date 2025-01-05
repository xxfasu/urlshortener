package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxfasu/urlshortener/internal/handler/v1/urls_handler"
	"github.com/xxfasu/urlshortener/internal/middleware"
)

func urlsRouter(router *gin.Engine, authM *middleware.AuthM, handler *urls_handler.Handler) {
	// URL缩短服务相关路由
	router.GET("/:code", handler.RedirectURL) // 短链接重定向

	// 需要JWT认证的URL管理API
	url := router.Group("/api")
	url.Use(authM.Handler())
	url.POST("/url", handler.CreateURL)                // 创建短链接
	url.GET("/urls", handler.GetURLs)                  // 获取用户的所有短链接
	url.DELETE("/url/:code", handler.DeleteURL)        // 删除短链接
	url.PATCH("/url/:code", handler.UpdateURLDuration) // 更新短链接的有效期

}
