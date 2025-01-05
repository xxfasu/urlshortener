package jwt

import (
	"fmt"
	"github.com/xxfasu/urlshortener/internal/conf"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type jwt struct {
}

func New() JWT {
	return &jwt{}
}

type UserCliams struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwtv5.RegisteredClaims
}

func (j *jwt) Generate(email string, useID int) (string, error) {
	claims := UserCliams{
		Email:  email,
		UserID: useID,
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(conf.Config.JWT.Expire)),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
		},
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(conf.Config.JWT.Secret)
}

func (j *jwt) ParseToken(tokenString string) (*UserCliams, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, &UserCliams{}, func(t *jwtv5.Token) (interface{}, error) {
		return conf.Config.JWT.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserCliams); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parseToken")
}
