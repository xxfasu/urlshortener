package repository

import (
	"github.com/google/wire"
	"github.com/xxfasu/urlshortener/internal/repository/urls_repository"
	"github.com/xxfasu/urlshortener/internal/repository/user_repository"
)

var ProviderSet = wire.NewSet(
	InitDB,
	NewTransaction,
	urls_repository.New,
	user_repository.New,
)
