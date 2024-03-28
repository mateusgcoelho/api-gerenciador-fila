package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	JwtPayload
	jwt.StandardClaims
}

type JwtPayload struct {
	Id          int `json:"id"`
	Permissions int `json:"permissions"`
}

func generateJwtToken(payload JwtPayload, secretKey string) (string, error) {
	claims := claims{
		JwtPayload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "api-gerenciador-fila",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
