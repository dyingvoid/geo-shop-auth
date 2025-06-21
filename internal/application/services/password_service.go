package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(bytes), nil
}

func (ps *PasswordService) Verify(hashed, provided string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(provided))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(bytes), nil
}
