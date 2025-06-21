package mocks

import (
	"geo-shop-auth/internal/domain"
	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) FindUserNickOrEmail(email, nickname string) (*domain.User, error) {
	args := m.Called(email, nickname)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserRepository) Insert(user *domain.User) (uuid.UUID, error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *UserRepository) FindUserNickname(nickname string) (*domain.User, error) {
	args := m.Called(nickname)
	return args.Get(0).(*domain.User), args.Error(1)
}
