package services

import "geo-shop-auth/internal/domain"

type TokenServicer interface {
	GenerateTokens() (domain.TokenPair, error)
	ParseAccessToken(str string) (*domain.AccessToken, error)
	FindRefreshToken(str string) (*domain.RefreshToken, error)
}

type PasswordServicer interface {
	Hash(password string) (string, error)
	Verify(hashed, provided string) error
}
