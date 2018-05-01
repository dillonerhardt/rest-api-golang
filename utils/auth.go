package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates a hash for the provided string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// VerifyPassword compares provided password string to provided hash
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
