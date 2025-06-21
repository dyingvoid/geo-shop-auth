package usecase_test

import (
	"geo-shop-auth/internal/application/usecase"
	"geo-shop-auth/internal/domain"
	"geo-shop-auth/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRefresh(t *testing.T) {
	mTokenService := new(mocks.TokenService)
	refreshToken := &domain.RefreshToken{
		Value:   uuid.UUID{},
		ExpTime: time.Now().Add(time.Hour).Unix(),
	}
	mTokenService.On("GenerateTokens").
		Return(domain.TokenPair{AccessToken: "", RefreshToken: refreshToken}, nil)
	mTokenService.On("FindRefreshToken", refreshToken.Value.String()).
		Return(refreshToken, nil)

	t.Run("successful refresh", func(t *testing.T) {
		req := usecase.RefreshRequest{
			RefreshToken: refreshToken.Value.String(),
		}
		_, err := usecase.Refresh(req, mTokenService)
		assert.NoError(t, err)

		if err != nil {
			panic(err)
		}
	})
}
