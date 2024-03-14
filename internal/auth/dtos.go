package auth

import "github.com/dgrijalva/jwt-go"

type JwtPayloadDto struct {
	Id         int `json:"id"`
	Permissoes int `json:"permissoes"`
	jwt.StandardClaims
}
