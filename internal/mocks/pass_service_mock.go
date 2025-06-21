package mocks

import "github.com/stretchr/testify/mock"

type PasswordService struct {
	mock.Mock
}

func (m *PasswordService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordService) Verify(hashed, provided string) error {
	args := m.Called(hashed, provided)
	return args.Error(0)
}
