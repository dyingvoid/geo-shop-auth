package repositories

import (
	"geo-shop-auth/internal/domain"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Insert(token *domain.RefreshToken) error
	FindToken(str string) (*domain.RefreshToken, error)
}

type UserRepository interface {
	Insert(u *domain.User) (uuid.UUID, error)
	FindUserNickname(nickname string) (*domain.User, error)
	FindUserNickOrEmail(email, nickname string) (*domain.User, error)
}
