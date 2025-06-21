package usecase_test

import (
	"errors"
	"geo-shop-auth/internal/application/common/error"
	"geo-shop-auth/internal/application/usecase"
	"geo-shop-auth/internal/domain"
	"geo-shop-auth/internal/mocks"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	t.Run("successful registration", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		mockRepo.On("FindUserNickOrEmail", "test@example.com", "testuser").
			Return((*domain.User)(nil), nil)
		id := uuid.New()
		mockRepo.On("Insert", mock.AnythingOfType("*domain.User")).
			Return(id, nil)

		req := usecase.RegisterRequest{
			Email:    "test@example.com",
			Password: "securepassword",
			Nickname: "testuser",
		}

		err := usecase.Register(req, mockRepo)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("validation fails - empty fields", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		req := usecase.RegisterRequest{
			Email:    "",
			Password: "",
			Nickname: "",
		}

		err := usecase.Register(req, mockRepo)

		assert.Error(t, err)
		var expectedErr *error.ValidationError
		assert.True(t, errors.As(err, &expectedErr))
		mockRepo.AssertNotCalled(t, "FindUserNickOrEmail")
	})

	t.Run("validation fails - invalid email", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		req := usecase.RegisterRequest{
			Email:    "invalid-email",
			Password: "password",
			Nickname: "user",
		}

		err := usecase.Register(req, mockRepo)

		assert.Error(t, err)
		var expectedErr *error.ValidationError
		assert.True(t, errors.As(err, &expectedErr))
		mockRepo.AssertNotCalled(t, "FindUserNickOrEmail")
	})

	t.Run("user already exists", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		existingEmail := "existing@example.com"
		existingNickname := "existingnickname"
		existingUser := &domain.User{}
		mockRepo.On("FindUserNickOrEmail", existingEmail, existingNickname).
			Return(existingUser, nil)

		req := usecase.RegisterRequest{
			Email:    existingEmail,
			Password: "pass",
			Nickname: existingNickname,
		}

		err := usecase.Register(req, mockRepo)

		assert.Error(t, err)
		var expectedErr *error.DuplicateError
		assert.True(t, errors.As(err, &expectedErr))
		mockRepo.AssertNotCalled(t, "Insert")
	})
}

func TestRegisterRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     usecase.RegisterRequest
		wantErr bool
	}{
		{
			name: "valid request",
			req: usecase.RegisterRequest{
				Email:    "valid@example.com",
				Password: "password",
				Nickname: "user",
			},
			wantErr: false,
		},
		{
			name: "invalid email format",
			req: usecase.RegisterRequest{
				Email:    "invalid-email",
				Password: "password",
				Nickname: "user",
			},
			wantErr: true,
		},
		{
			name: "empty password",
			req: usecase.RegisterRequest{
				Email:    "valid@example.com",
				Password: "",
				Nickname: "user",
			},
			wantErr: true,
		},
		{
			name: "empty nickname",
			req: usecase.RegisterRequest{
				Email:    "valid@example.com",
				Password: "password",
				Nickname: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
