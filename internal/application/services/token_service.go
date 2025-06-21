package services

import (
	"fmt"
	"geo-shop-auth/internal/application/common"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type TokenService struct {
	rep     repositories.TokenRepository
	options common.JWTOptions
}

func NewTokenService(
	rep repositories.TokenRepository,
	options common.JWTOptions,
) *TokenService {
	return &TokenService{
		rep:     rep,
		options: options,
	}
}

func (ts *TokenService) GenerateTokens() (domain.TokenPair, error) {
	accessToken, err := ts.generateAccessToken()
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken := ts.generateRefreshToken()
	err = ts.rep.Insert(refreshToken)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("error storing refresh token: %w", err)
	}

	return domain.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (ts *TokenService) ParseAccessToken(
	tokenString string,
) (*domain.AccessToken, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&domain.AccTknClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != ts.options.SigningMethod {
				return nil, fmt.Errorf(
					"unexpexted signing method: %v (expected %v)",
					token.Header["alg"],
					ts.options.SigningMethod.Alg(),
				)
			}

			return ts.options.SigningKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(*domain.AccTknClaims); ok {
		return &domain.AccessToken{
			Token:  token,
			Claims: claims,
		}, nil
	}

	return nil, fmt.Errorf("bad claims")
}

func (ts *TokenService) FindRefreshToken(
	tokenString string,
) (*domain.RefreshToken, error) {
	tkn, err := ts.rep.FindToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("error fetching token: %w", err)
	}

	return tkn, nil
}

func (ts *TokenService) generateAccessToken() (string, error) {
	claims := domain.AccTknClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ts.options.AccessTknDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-service",
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(ts.options.SigningMethod, claims)
	signedToken, err := token.SignedString(ts.options.SigningKey)
	if err != nil {
		return "", fmt.Errorf("error signing access token: %w", err)
	}

	return signedToken, nil
}

func (ts *TokenService) generateRefreshToken() *domain.RefreshToken {
	return &domain.RefreshToken{
		Value:   uuid.New(),
		ExpTime: time.Now().Add(ts.options.RefreshTknDuration).Unix(),
	}
}
