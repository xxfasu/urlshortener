//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/middleware"
	"github.com/xxfasu/urlshortener/routes"
)

func NewWire() (*gin.Engine, func(), error) {
	panic(wire.Build(
		middleware.ProviderSet,
		routes.ProviderSet,
	))
}
