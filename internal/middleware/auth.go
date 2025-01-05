package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xxfasu/urlshortener/pkg/jwt"
	"net/http"
	"strings"
)

type AuthM struct {
	j jwt.JWT
}

func NewAuthM(j jwt.JWT) *AuthM {
	return &AuthM{

		j: j,
	}
}

func (m *AuthM) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		ls := strings.Split(authHeader, " ")

		if len(ls) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization format is incorrect"})
			return
		}
		if ls[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization format is incorrect"})
			return
		}
		tokenString := ls[1]

		claims, err := m.j.ParseToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "jwt parse error"})
			return
		}
		ctx.Set("email", claims.Email)
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
