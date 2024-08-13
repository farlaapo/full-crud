package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, HashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
	return err == nil
}
