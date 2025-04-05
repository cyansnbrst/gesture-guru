package auth

import "context"

// Auth usecase interface
type UseCase interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	Register(ctx context.Context, email string, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
