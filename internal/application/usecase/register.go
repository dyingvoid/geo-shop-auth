package usecase

import (
	"fmt"
	"geo-shop-auth/internal/application/common/error"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func Register(
	req RegisterRequest,
	rep repositories.UserRepository,
) error {
	err := req.Validate()
	if err != nil {
		return fmt.Errorf("error validating request: %w", err)
	}

	user, err := rep.FindUserNickOrEmail(
		req.Email,
		req.Nickname,
	)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	if user != nil {
		return &error.DuplicateError{
			Msg: "user already exists",
		}
	}

	hashed, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = rep.Insert(domain.NewUser(req.Email, req.Nickname, hashed))
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
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
		return &error.ValidationError{
			Msg: "password or nickname is empty",
		}
	}
	if len(r.Email) > 0 {
		_, err := mail.ParseAddress(r.Email)
		if err != nil {
			return &error.ValidationError{
				Msg: fmt.Sprintf("invalid email address: %s, %v", r.Email, err),
			}
		}
	}

	return nil
}
