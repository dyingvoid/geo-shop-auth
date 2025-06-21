package mocks

import (
	"geo-shop-auth/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TokenRepository struct {
	mock.Mock
}

func (m *TokenRepository) Insert(token *domain.RefreshToken) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *TokenRepository) FindToken(token string) (*domain.RefreshToken, error) {
	args := m.Called(token)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}
