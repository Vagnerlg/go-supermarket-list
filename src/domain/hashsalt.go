package domain

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashSalt(secret string) (string, error) {
	hashsalt, erro := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)

	return string(hashsalt), erro
}
