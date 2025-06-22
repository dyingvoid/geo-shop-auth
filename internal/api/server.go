package api

import (
	"context"
	pb "geo-shop-auth/internal/api/gen/authpb"
	"geo-shop-auth/internal/application/repositories"
	"geo-shop-auth/internal/application/services"
	"geo-shop-auth/internal/application/usecase"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	userRepository  repositories.UserRepository
	tokenService    services.TokenServicer
	passwordService services.PasswordServicer
}

func NewServer(
	userRepository repositories.UserRepository,
	tokenService services.TokenServicer,
	passwordService services.PasswordServicer,
) *Server {
	return &Server{
		userRepository:  userRepository,
		tokenService:    tokenService,
		passwordService: passwordService,
	}
}

func (s *Server) Register(
	ctx context.Context,
	in *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	req := &usecase.RegisterRequest{
		Email:    in.Email,
		Password: in.Password,
		Nickname: in.Nickname,
	}
	err := usecase.Register(ctx, req, s.userRepository)

	return &pb.RegisterResponse{}, err
}

func (s *Server) Login(
	ctx context.Context,
	in *pb.LoginRequest,
) (*pb.LoginResponse, error) {
	req := &usecase.LoginRequest{
		Nickname: in.Nickname,
		Password: in.Password,
	}
	res, err := usecase.Login(ctx, req, s.userRepository, s.tokenService, s.passwordService)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, err
}

func (s *Server) Refresh(
	ctx context.Context,
	in *pb.RefreshRequest,
) (*pb.RefreshResponse, error) {
	req := &usecase.RefreshRequest{
		RefreshToken: in.RefreshToken,
	}
	res, err := usecase.Refresh(ctx, req, s.tokenService)
	if err != nil {
		return nil, err
	}
	return &pb.RefreshResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
