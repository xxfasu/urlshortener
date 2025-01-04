package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/conf"
	"github.com/xxfasu/urlshortener/internal/middleware"
	"net/http"
)

var ProviderSet = wire.NewSet(NewRouter)

func NewRouter(
	recoveryM *middleware.Recovery,
	corsM *middleware.Cors,
	logM *middleware.LogM,
) *gin.Engine {
	router := gin.New()
	if conf.Env.Environment == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(recoveryM.Handler())
	router.Use(corsM.Handler())
	router.Use(logM.Handler())

	{
		// 健康监测
		router.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

	}
	return router
}
