//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/handler"
	"github.com/xxfasu/urlshortener/internal/middleware"
	"github.com/xxfasu/urlshortener/internal/repository"
	"github.com/xxfasu/urlshortener/internal/service"
	"github.com/xxfasu/urlshortener/pkg/jwt"
	"github.com/xxfasu/urlshortener/routes"
)

func NewWire() (*gin.Engine, func(), error) {
	panic(wire.Build(
		jwt.New,
		middleware.ProviderSet,
		routes.ProviderSet,
		repository.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
	))
}
