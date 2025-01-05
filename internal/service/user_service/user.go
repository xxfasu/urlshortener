package user_service

import (
	"context"
	"github.com/xxfasu/urlshortener/internal/common/response"
	"github.com/xxfasu/urlshortener/internal/repository"
	"github.com/xxfasu/urlshortener/internal/repository/user_repository"
	"github.com/xxfasu/urlshortener/internal/validation"
)

type service struct {
	tm        repository.Transaction
	usersRepo user_repository.Repository
}

func New(
	tm repository.Transaction,
	usersRepo user_repository.Repository,
) Service {
	return &service{
		tm:        tm,
		usersRepo: usersRepo,
	}
}

func (s *service) Login(ctx context.Context, req *validation.Login) (*response.Login, error) {

}

func (s *service) IsEmailAvailable(ctx context.Context, email string) error {

	return nil
}

func (s *service) Register(ctx context.Context, req *validation.Register) (*response.Login, error) {

}

func (s *service) SendEmailCode(ctx context.Context, email string) error {

	return nil
}

func (s *service) ResetPassword(ctx context.Context, req *validation.ForgetPassword) (*response.Login, error) {

}
