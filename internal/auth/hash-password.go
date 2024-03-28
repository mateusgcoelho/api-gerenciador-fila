package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePasswords(hashedPassword string, plainText string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText)) == nil
}
