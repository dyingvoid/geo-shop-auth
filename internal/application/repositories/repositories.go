package repositories

import (
	"context"
	"geo-shop-auth/internal/domain"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Insert(ctx context.Context, token *domain.RefreshToken) error
	FindToken(ctx context.Context, str string) (*domain.RefreshToken, error)
}

type UserRepository interface {
	Insert(ctx context.Context, u *domain.User) (uuid.UUID, error)
	FindUserNickname(ctx context.Context, nickname string) (*domain.User, error)
	FindUserNickOrEmail(ctx context.Context, email, nickname string) (*domain.User, error)
}
