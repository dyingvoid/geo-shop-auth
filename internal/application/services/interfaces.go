package services

import (
	"context"
	"geo-shop-auth/internal/domain"
)

type TokenServicer interface {
	GenerateTokens(ctx context.Context) (domain.TokenPair, error)
	ParseAccessToken(ctx context.Context, str string) (*domain.AccessToken, error)
	FindRefreshToken(ctx context.Context, str string) (*domain.RefreshToken, error)
}

type PasswordServicer interface {
	Hash(password string) (string, error)
	Verify(hashed, provided string) error
}
