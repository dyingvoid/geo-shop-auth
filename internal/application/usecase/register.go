package usecase

import (
	"context"
	"fmt"
	"geo-shop-auth/internal/application/common/commonerror"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func Register(
	ctx context.Context,
	req *RegisterRequest,
	rep repositories.UserRepository,
) error {
	err := req.Validate()
	if err != nil {
		return fmt.Errorf("commonerror validating request: %w", err)
	}

	user, err := rep.FindUserNickOrEmail(
		ctx,
		req.Email,
		req.Nickname,
	)
	if err != nil {
		return fmt.Errorf("commonerror fetching user: %w", err)
	}
	if user != nil {
		return &commonerror.DuplicateError{
			Msg: "user already exists",
		}
	}

	hashed, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = rep.Insert(ctx, domain.NewUser(req.Email, req.Nickname, hashed))
	if err != nil {
		return fmt.Errorf("commonerror inserting user: %w", err)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("commonerror hashing password: %w", err)
	}

	return string(bytes), nil
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func (r *RegisterRequest) Validate() error {
	if r.Password == "" || r.Nickname == "" {
		return &commonerror.ValidationError{
			Msg: "password or nickname is empty",
		}
	}
	if len(r.Email) > 0 {
		_, err := mail.ParseAddress(r.Email)
		if err != nil {
			return &commonerror.ValidationError{
				Msg: fmt.Sprintf("invalid email address: %s, %v", r.Email, err),
			}
		}
	}

	return nil
}
