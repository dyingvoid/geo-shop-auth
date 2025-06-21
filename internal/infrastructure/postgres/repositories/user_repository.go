package repositories

import (
	"context"
	"database/sql"
	"geo-shop-auth/internal/domain"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Insert(ctx context.Context, u *domain.User) (uuid.UUID, error) {
	_, err := ur.db.ExecContext(ctx,
		`INSERT INTO users (id, email, nickname, pass_hash) VALUES ($1, $2, $3, $4)`,
		u.ID, u.Email, u.Nickname, u.PassHash,
	)
	if err != nil {
		return uuid.Nil, err
	}

	return u.ID, nil
}

func (ur *UserRepository) FindUserNickname(ctx context.Context, nickname string) (*domain.User, error) {
	var u domain.User
	row := ur.db.QueryRowContext(ctx,
		`SELECT * FROM users WHERE nickname = $1`,
		nickname,
	)
	err := userRowScan(row, &u)
	return &u, err
}

func (ur *UserRepository) FindUserNickOrEmail(
	ctx context.Context,
	email,
	nickname string,
) (*domain.User, error) {
	var u domain.User
	row := ur.db.QueryRowContext(ctx,
		`SELECT * FROM users WHERE email = $1 OR nickname = $2`,
		email, nickname,
	)
	err := userRowScan(row, &u)
	return &u, err
}

func userRowScan(row *sql.Row, u *domain.User) error {
	err := row.Scan(&u.ID, &u.Email, &u.Nickname, &u.PassHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}
