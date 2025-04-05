package usecase

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/cyansnbrst/gesture-guru/sso-service/config"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/db"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/jwt"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/logger"
)

type authUC struct {
	logger   *zap.Logger
	cfg      *config.Config
	authRepo auth.Repository
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewAuthUseCase(logger *zap.Logger, cfg *config.Config, authRepo auth.Repository) auth.UseCase {
	return &authUC{logger: logger, cfg: cfg, authRepo: authRepo}
}

// Login checks if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (uc *authUC) Login(ctx context.Context, email string, password string) (string, error) {
	const op = "Auth.Login"

	user, err := uc.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.LogError(uc.logger, op, err, "failed to get user")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	token, err := jwt.NewToken(user, uc.cfg)
	if err != nil {
		logger.LogError(uc.logger, op, err, "failed to generate token")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (uc *authUC) Register(ctx context.Context, email string, password string) (userID int64, err error) {
	const op = "Auth.Register"

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.LogError(uc.logger, op, err, "failed to generate password hash")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := uc.authRepo.SaveUser(ctx, email, passHash)
	if err != nil {
		if !errors.Is(err, db.ErrUserExists) {
			logger.LogError(uc.logger, op, err, "failed to save user")
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (uc *authUC) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "Auth.IsAdmin"

	isAdmin, err := uc.authRepo.IsAdmin(ctx, userID)
	if err != nil {
		logger.LogError(uc.logger, op, err, "failed to check if user is admin")
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
