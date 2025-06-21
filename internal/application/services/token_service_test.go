package services_test

import (
	"crypto/rand"
	"geo-shop-auth/internal/application/common"
	"geo-shop-auth/internal/application/services"
	"geo-shop-auth/internal/domain"
	"geo-shop-auth/internal/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestTokenService(t *testing.T) {
	t.Run("token valid", func(t *testing.T) {
		repoMock := new(mocks.TokenRepository)
		repoMock.
			On("Insert", mock.Anything, mock.Anything).
			Return(nil)
		repoMock.On("FindToken", mock.Anything).
			Return(&domain.RefreshToken{}, nil)
		jwtOptions := common.JWTOptions{
			AccessTknDuration:  time.Hour,
			RefreshTknDuration: time.Hour * 2,
			SigningMethod:      jwt.SigningMethodHS256,
			SigningKey:         generateHS256Key(),
		}
		tokenService := services.NewTokenService(repoMock, jwtOptions)

		tokenPair, err := tokenService.GenerateTokens()
		assert.NoError(t, err)

		accTkn, err := tokenService.ParseAccessToken(tokenPair.AccessToken)
		assert.NoError(t, err)
		assert.True(t, accTkn.Valid)
		assert.NoError(t, err)

		_, err = tokenService.FindRefreshToken(tokenPair.RefreshToken.String())
		assert.NoError(t, err)
		repoMock.AssertExpectations(t)
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
