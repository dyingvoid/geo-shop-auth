package usecase

import (
	"fmt"
	"geo-shop-auth/internal/application/services"
	"time"
)

func Refresh(
	req RefreshRequest,
	tokenService services.TokenServicer,
) (RefreshResponse, error) {
	token, err := tokenService.FindRefreshToken(req.RefreshToken)
	if err != nil {
		return RefreshResponse{}, fmt.Errorf("error parsing token: %w", err)
	}

	if token.ExpTime < time.Now().Unix() {
		return RefreshResponse{}, fmt.Errorf("token expired")
	}

	tokenPair, err := tokenService.GenerateTokens()
	if err != nil {
		return RefreshResponse{}, fmt.Errorf("error generating tokens: %w", err)
	}

	return RefreshResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken.String(),
	}, nil
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
