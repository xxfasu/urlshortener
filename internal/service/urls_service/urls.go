package urls_service

import (
	"context"
	"github.com/xxfasu/urlshortener/internal/common/response"
	"github.com/xxfasu/urlshortener/internal/repository"
	"github.com/xxfasu/urlshortener/internal/repository/urls_repository"
	"github.com/xxfasu/urlshortener/internal/validation"
)

type service struct {
	tm       repository.Transaction
	urlsRepo urls_repository.Repository
}

func New(
	tm repository.Transaction,
	urlsRepo urls_repository.Repository,
) Service {
	return &service{
		tm:       tm,
		urlsRepo: urlsRepo,
	}
}

func (s *service) CreateURL(ctx context.Context, req *validation.CreateURL) (shortURL string, err error) {

}

func (s *service) GetURL(ctx context.Context, shortCode string) (originalURL string, err error) {

}

func (s *service) IncrViews(ctx context.Context, shortCode string) error {
	return nil
}

func (s *service) GetURLs(ctx context.Context, req *validation.GetURLs) (*response.GetURLs, error) {

}

func (s *service) DeleteURL(ctx context.Context, shortCode string) error {

	return nil
}

func (s *service) UpdateURLDuration(ctx context.Context, req *validation.UpdateURLDuration) error {
	return nil
}
