package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(origin string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(origin), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}

func ComparePassword(provided, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(provided))
	if err != nil {
		// This error specifically indicates a mismatch.
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("invalid password")
		}
		// Other errors could indicate issues with the hash itself.
		return errors.New("error comparing password")
	}
	return nil // Passwords match
}
