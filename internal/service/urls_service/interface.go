package urls_service

import (
	"context"
	"github.com/xxfasu/urlshortener/internal/common/response"
	"github.com/xxfasu/urlshortener/internal/validation"
)

type Service interface {
	CreateURL(ctx context.Context, req *validation.CreateURL) (shortURL string, err error)
	GetURL(ctx context.Context, shortCode string) (originalURL string, err error)
	IncrViews(ctx context.Context, shortCode string) error
	GetURLs(ctx context.Context, req *validation.GetURLs) (*response.GetURLs, error)
	DeleteURL(ctx context.Context, shortCode string) error
	UpdateURLDuration(ctx context.Context, req *validation.UpdateURLDuration) error
}
