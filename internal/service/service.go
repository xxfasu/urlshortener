package service

import (
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/service/urls_service"

	"github.com/xxfasu/urlshortener/internal/service/user_service"
)

var ProviderSet = wire.NewSet(
	urls_service.New,
	user_service.New,
)
