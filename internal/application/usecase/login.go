package usecase

import (
	"fmt"
	"geo-shop-auth/internal/application/common/error"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/application/services"
)

func Login(
	req LoginRequest,
	userRep repositories.UserRepository,
	tokenService services.TokenServicer,
	passwordService services.PasswordServicer,
) (LoginResponse, error) {
	err := req.Validate()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error validating login request: %w", err)
	}

	usr, err := userRep.FindUserNickname(req.Nickname)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error fetching user data: %w", err)
	}
	if usr == nil {
		return LoginResponse{}, fmt.Errorf("invalid user or password")
	}

	err = passwordService.Verify(usr.PassHash, req.Password)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("invalid user or password")
	}

	tokenPair, err := tokenService.GenerateTokens()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("error generating tokens: %w", err)
	}

	return LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken.String(),
	}, nil
}

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if r.Nickname == "" || r.Password == "" {
		return &error.ValidationError{
			Msg: "nickname and password are required",
		}
	}

	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
