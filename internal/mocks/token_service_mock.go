package mocks

import (
	"geo-shop-auth/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TokenService struct {
	mock.Mock
}

func (m *TokenService) GenerateTokens() (domain.TokenPair, error) {
	args := m.Called()
	return args.Get(0).(domain.TokenPair), args.Error(1)
}

func (m *TokenService) ParseAccessToken(str string) (*domain.AccessToken, error) {
	args := m.Called(str)
	return args.Get(0).(*domain.AccessToken), args.Error(1)
}

func (m *TokenService) FindRefreshToken(str string) (*domain.RefreshToken, error) {
	args := m.Called(str)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}
