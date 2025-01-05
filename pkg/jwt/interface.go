package jwt

type JWT interface {
	Generate(email string, useID int) (string, error)
	ParseToken(tokenString string) (*UserCliams, error)
}
