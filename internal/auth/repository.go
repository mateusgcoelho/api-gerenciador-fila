package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	GenerateJwtToken(payload JwtPayloadDto) (string, error)
	HashPassword(password string) (string, error)
	ValidateToken(bearerToken string) (JwtPayloadDto, error)
	ComparePasswords(hashedPassword, attemptedPassword string) bool
}

type authRepositoryImpl struct {
}

func NewAuthRepository() IAuthRepository {
	return authRepositoryImpl{}
}

func (r authRepositoryImpl) GenerateJwtToken(payload JwtPayloadDto) (string, error) {
	claims := JwtPayloadDto{
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

func (r authRepositoryImpl) HashPassword(password string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return "", errors.New("Não foi possível criptografar senha.")
	}

	return string(passwordHashed), nil
}

func (r authRepositoryImpl) ComparePasswords(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func (r authRepositoryImpl) ValidateToken(bearerToken string) (JwtPayloadDto, error) {
	token, err := jwt.ParseWithClaims(bearerToken, &JwtPayloadDto{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return JwtPayloadDto{}, err
	}

	payload, ok := token.Claims.(*JwtPayloadDto)
	if !ok || !token.Valid {
		return JwtPayloadDto{}, errors.New("Token de autenticação inválido.")
	}

	return *payload, nil
}
