package handler

import (
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/handler/v1/urls_handler"
	"github.com/xxfasu/urlshortener/internal/handler/v1/user_handler"
)

var ProviderSet = wire.NewSet(
	urls_handler.New,
	user_handler.New,
)
