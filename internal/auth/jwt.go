package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Id         int `json:"id"`
	Permissoes int `json:"permissoes"`
	jwt.StandardClaims
}

type JwtPayload struct {
	Id         int
	Permissoes int
}

func generateJwtToken(payload JwtPayload) (string, error) {
	claims := claims{
		Id:         payload.Id,
		Permissoes: payload.Permissoes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "api-gerenciador-fila",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
