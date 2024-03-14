package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	salt := generateSalt(8)

	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func generateSalt(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return strings.ReplaceAll(fmt.Sprintf("%x", bytes), " ", "")
}

func ComparePasswords(hashedPassword, attemptedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(attemptedPassword)) == nil
}
