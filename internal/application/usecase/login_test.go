package usecase_test

import (
	"crypto/rand"
	"errors"
	"fmt"
	"geo-shop-auth/internal/application/common/commonerror"
	"geo-shop-auth/internal/application/usecase"
	"geo-shop-auth/internal/domain"
	"geo-shop-auth/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	mTokenService := new(mocks.TokenService)
	refreshToken := &domain.RefreshToken{Value: uuid.UUID{}, ExpTime: time.Now().Add(time.Hour).Unix()}
	mTokenService.On("GenerateTokens").
		Return(domain.TokenPair{AccessToken: "", RefreshToken: refreshToken}, nil)

	mPasswordService := new(mocks.PasswordService)
	mPasswordService.On("Verify", "goodpass", "goodpass").
		Return(error(nil))
	mPasswordService.On("Verify", "goodpass", "badpass").
		Return(fmt.Errorf("invalid pass"))

	mUserRep := new(mocks.UserRepository)
	usr := domain.NewUser("email", "nickname", "goodpass")
	mUserRep.On("FindUserNickname", "exist").
		Return(usr, nil)
	mUserRep.On("FindUserNickname", "nil").
		Return((*domain.User)(nil), fmt.Errorf("user not found"))

	t.Run("successful login", func(t *testing.T) {
		req := usecase.LoginRequest{
			Nickname: "exist",
			Password: "goodpass",
		}
		_, err := usecase.Login(req, mUserRep, mTokenService, mPasswordService)

		assert.NoError(t, err)
	})

	t.Run("validation fails - empty req fiels", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		tokenRepMock := new(mocks.TokenRepository)
		tokenRepMock.
			On("Insert", mock.Anything, mock.Anything).
			Return(nil)
		req := usecase.LoginRequest{
			Nickname: "",
			Password: "",
		}
		_, err := usecase.Login(req, mockRepo, mTokenService, mPasswordService)
		assert.Error(t, err)
		var valErr *commonerror.ValidationError
		assert.True(t, errors.As(err, &valErr))
	})

	t.Run("login fails - user does not exist", func(t *testing.T) {
		req := usecase.LoginRequest{
			Nickname: "nil",
			Password: "password",
		}
		_, err := usecase.Login(req, mUserRep, mTokenService, mPasswordService)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

func generateHS256Key() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}
