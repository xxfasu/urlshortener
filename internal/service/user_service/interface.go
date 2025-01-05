package user_service

import (
	"context"
	"github.com/xxfasu/urlshortener/internal/common/response"
	"github.com/xxfasu/urlshortener/internal/validation"
)

type Service interface {
	Login(ctx context.Context, req *validation.Login) (*response.Login, error)
	IsEmailAvailable(ctx context.Context, email string) error
	Register(ctx context.Context, req *validation.Register) (*response.Login, error)
	SendEmailCode(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, req *validation.ForgetPassword) (*response.Login, error)
}
