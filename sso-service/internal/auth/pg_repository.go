package auth

import (
	"context"

	"github.com/cyansnbrst/gesture-guru/sso-service/internal/models"
)

// Auth repository interface
type Repository interface {
	SaveUser(ctx context.Context, email string, passwordHash []byte) (uid int64, err error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
