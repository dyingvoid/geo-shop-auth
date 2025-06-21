package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessToken struct {
	*jwt.Token
	Claims *AccTknClaims
}

type AccTknClaims struct {
	jwt.RegisteredClaims
}
