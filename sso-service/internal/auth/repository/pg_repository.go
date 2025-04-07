package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/models"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/db"
)

// Auth repository struct
type authRepo struct {
	db *pgxpool.Pool
}

// Auth repository constructor
func NewAuthRepo(db *pgxpool.Pool) auth.Repository {
	return &authRepo{db: db}
}

// SaveUser saves a new user to db.
func (r *authRepo) SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	const op = "repository.SaveUser"

	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(ctx, query, email, passwordHash).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, db.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// GerUserByEmail returns user by email.
func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	const op = "repository.GetUserByEmail"

	query := `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, db.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

// IsAdmin checks if user is admin.
func (r *authRepo) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "repository.IsAdmin"

	query := `
		SELECT is_admin 
		FROM users 
		WHERE id = $1
	`

	var isAdmin bool
	err := r.db.QueryRow(ctx, query, userID).Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, db.ErrUserNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
