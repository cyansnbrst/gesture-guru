package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"cyansnbrst/gestures-service/internal/gestures"
	"cyansnbrst/gestures-service/internal/models"
	"cyansnbrst/gestures-service/pkg/db"
)

// Gestures repository struct
type gesturesRepo struct {
	db *pgxpool.Pool
}

// New gestures repository
func NewGesturesRepo(db *pgxpool.Pool) gestures.Repository {
	return &gesturesRepo{db: db}
}

// Get gesture by ID
func (r *gesturesRepo) GetByID(ctx context.Context, id int64) (*models.Gesture, error) {
	query := `
		SELECT id, name, description, video_url, category_id, created_at
		FROM gestures
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var gesture models.Gesture
	err := row.Scan(&gesture.ID, &gesture.Name, &gesture.Description, &gesture.VideoURL, &gesture.CategoryID, &gesture.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, db.ErrGestureNotFound
		}
		return nil, fmt.Errorf("repo - failed to scan row: %w", err)
	}

	return &gesture, nil
}

// Get all gestures
func (r *gesturesRepo) GetAll(ctx context.Context) ([]models.Gesture, error) {
	query := `
		SELECT id, name, description, video_url, category_id, created_at
		FROM gestures
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repo - failed to execute query: %w", err)
	}
	defer rows.Close()

	var gestures []models.Gesture
	for rows.Next() {
		var gesture models.Gesture
		if err := rows.Scan(&gesture.ID, &gesture.Name, &gesture.Description, &gesture.VideoURL, &gesture.CategoryID, &gesture.CreatedAt); err != nil {
			return nil, fmt.Errorf("repo - failed to scan row: %w", err)
		}
		gestures = append(gestures, gesture)
	}

	if len(gestures) == 0 {
		return nil, db.ErrGestureNotFound
	}

	return gestures, nil
}

// Create a new gesture
func (r *gesturesRepo) Create(ctx context.Context, gesture *models.Gesture) (int64, error) {
	query := `
		INSERT INTO gestures (name, description, video_url, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(ctx, query, gesture.Name, gesture.Description, gesture.VideoURL, gesture.CategoryID).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return 0, db.ErrInvalidCategory
			}
		}
		return 0, fmt.Errorf("repo - failed to insert gesture: %w", err)
	}

	return id, nil
}

// Update gesture by ID
func (r *gesturesRepo) Update(ctx context.Context, gesture *models.Gesture) error {
	query := `
		UPDATE gestures
		SET name = $1, description = $2, video_url = $3, category_id = $4
		WHERE id = $5
		RETURNING id
	`

	var updatedID int64
	err := r.db.QueryRow(ctx, query, gesture.Name, gesture.Description, gesture.VideoURL, gesture.CategoryID, gesture.ID).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.ErrGestureNotFound
		}
		return fmt.Errorf("repo - failed to update gesture: %w", err)
	}

	return nil
}

// Delete gesture by ID
func (r *gesturesRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM gestures WHERE id = $1 RETURNING id`

	var deletedID int64
	err := r.db.QueryRow(ctx, query, id).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.ErrGestureNotFound
		}
		return fmt.Errorf("repo - failed to delete gesture: %w", err)
	}

	return nil
}
