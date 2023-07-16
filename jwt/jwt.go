package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Uid string
	jwt.RegisteredClaims
}

type Jwt struct {
	Secret string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{Secret: secret}
}

func (j Jwt) CreateToken(claims Claims) (string, error) {
	accessJwtKey := []byte(j.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(accessJwtKey)
}

func (Jwt) GetPayload(token string) (string, error) {
	parser := jwt.NewParser()
	var claims Claims
	_, _, err := parser.ParseUnverified(token, &claims)
	return claims.Uid, err
}

func (j Jwt) VerifyToken(token string) error {
	accessJwtKey := []byte(j.Secret)
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return accessJwtKey, nil
	})
	return err
}
