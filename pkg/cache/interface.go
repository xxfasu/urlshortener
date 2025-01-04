package cache

import (
	"context"
	"time"
)

type Cache interface {
	GetCacheStatus() StatusInfo
	GetKeyStatus(key string) KeyInfo
	GetCache(ctx context.Context, key string) (string, error)
	GetLocal(key string) (string, error)
	GetCacheOrElse(ctx context.Context, key string, ttl time.Duration, fetcher Fetcher) (string, error)
	SetCache(ctx context.Context, key, value string, ttl time.Duration) error
	SetLocal(key, value string, ttl time.Duration) error
	DelCache(key string) error
	DelLocal(key string) error
}

type Fetcher interface {
	Fetch(ctx context.Context, key string) (string, error)
}
