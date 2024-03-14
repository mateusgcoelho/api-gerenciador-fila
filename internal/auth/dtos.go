package auth

import "github.com/dgrijalva/jwt-go"

type claims struct {
	Id         int `json:"id"`
	Permissoes int `json:"permissoes"`
	jwt.StandardClaims
}

type JwtPayloadDto struct {
	Id         int
	Permissoes int
}
