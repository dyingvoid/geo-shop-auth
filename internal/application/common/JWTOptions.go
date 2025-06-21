package common

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTOptions struct {
	AccessTknDuration  time.Duration
	RefreshTknDuration time.Duration
	SigningMethod      jwt.SigningMethod
	SigningKey         []byte
}

func (o JWTOptions) AccTknExpTime() int64 {
	return time.Now().Add(o.AccessTknDuration).Unix()
}

func (o JWTOptions) RefreshTknExpTime() int64 {
	return time.Now().Add(o.RefreshTknDuration).Unix()
}
