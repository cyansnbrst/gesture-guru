package grpc

import (
	"context"
	"errors"

	ssov1 "github.com/cyansnbrst/gesture-guru/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth/usecase"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/db"
)

// Auth server struct
type authServer struct {
	ssov1.UnimplementedAuthServer
	authUC auth.UseCase
}

// Auth server constructor
func NewAuthServer(gRPCServer *grpc.Server, authUC auth.UseCase) {
	ssov1.RegisterAuthServer(gRPCServer, &authServer{authUC: authUC})
}

// Login user
func (s *authServer) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if in.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	token, err := s.authUC.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

// Register user
func (s *authServer) Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.authUC.Register(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, db.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{UserId: uid}, nil
}

// Check user's admin status
func (s *authServer) IsAdmin(ctx context.Context, in *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	isAdmin, err := s.authUC.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "failed to check admin status")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}
