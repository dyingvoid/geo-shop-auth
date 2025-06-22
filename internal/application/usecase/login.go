package usecase

import (
	"context"
	"fmt"
	"geo-shop-auth/internal/application/common/commonerror"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/application/services"
)

func Login(
	ctx context.Context,
	req *LoginRequest,
	userRep repositories.UserRepository,
	tokenService services.TokenServicer,
	passwordService services.PasswordServicer,
) (LoginResponse, error) {
	err := req.Validate()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("commonerror validating login request: %w", err)
	}

	usr, err := userRep.FindUserNickname(ctx, req.Nickname)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("commonerror fetching user data: %w", err)
	}
	if usr == nil {
		return LoginResponse{}, fmt.Errorf("invalid user or password")
	}

	err = passwordService.Verify(usr.PassHash, req.Password)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("invalid user or password")
	}

	tokenPair, err := tokenService.GenerateTokens(ctx)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("commonerror generating tokens: %w", err)
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
		return &commonerror.ValidationError{
			Msg: "nickname and password are required",
		}
	}

	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
