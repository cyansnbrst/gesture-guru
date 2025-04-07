package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
	"github.com/cyansnbrst/gesture-guru/gestures-service/pkg/db"
)

type gesturesRepo struct {
	db *pgxpool.Pool
}

func NewGesturesRepo(db *pgxpool.Pool) gestures.Repository {
	return &gesturesRepo{db: db}
}

// GetByID returns a gesture by its ID.
func (r *gesturesRepo) GetByID(ctx context.Context, id int64) (*models.Gesture, error) {
	const op = "repository.GetByID"

	query := `
		SELECT id, name, description, video_url, category_id, created_at
		FROM gestures
		WHERE id = $1
	`

	var gesture models.Gesture
	err := r.db.QueryRow(ctx, query, id).Scan(
		&gesture.ID,
		&gesture.Name,
		&gesture.Description,
		&gesture.VideoURL,
		&gesture.CategoryID,
		&gesture.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, db.ErrGestureNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &gesture, nil
}

// GetAll returns all gestures.
func (r *gesturesRepo) GetAll(ctx context.Context) ([]models.Gesture, error) {
	const op = "repository.GetAll"

	query := `
		SELECT id, name, description, video_url, category_id, created_at
		FROM gestures
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var gestures []models.Gesture
	for rows.Next() {
		var gesture models.Gesture
		if err := rows.Scan(
			&gesture.ID,
			&gesture.Name,
			&gesture.Description,
			&gesture.VideoURL,
			&gesture.CategoryID,
			&gesture.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		gestures = append(gestures, gesture)
	}

	if len(gestures) == 0 {
		return nil, fmt.Errorf("%s: %w", op, db.ErrGestureNotFound)
	}

	return gestures, nil
}

// Create inserts a new gesture and returns its ID.
func (r *gesturesRepo) Create(ctx context.Context, gesture *models.Gesture) (int64, error) {
	const op = "repository.Create"

	query := `
		INSERT INTO gestures (name, description, video_url, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(
		ctx,
		query,
		gesture.Name,
		gesture.Description,
		gesture.VideoURL,
		gesture.CategoryID,
	).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return 0, fmt.Errorf("%s: %w", op, db.ErrInvalidCategory)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// Update modifies a gesture by ID.
func (r *gesturesRepo) Update(ctx context.Context, gesture *models.Gesture) error {
	const op = "repository.Update"

	query := `
		UPDATE gestures
		SET name = $1, description = $2, video_url = $3, category_id = $4
		WHERE id = $5
		RETURNING id
	`

	var updatedID int64
	err := r.db.QueryRow(
		ctx,
		query,
		gesture.Name,
		gesture.Description,
		gesture.VideoURL,
		gesture.CategoryID,
		gesture.ID,
	).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, db.ErrGestureNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Delete removes a gesture by ID.
func (r *gesturesRepo) Delete(ctx context.Context, id int64) error {
	const op = "repository.Delete"

	query := `DELETE FROM gestures WHERE id = $1 RETURNING id`

	var deletedID int64
	err := r.db.QueryRow(ctx, query, id).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, db.ErrGestureNotFound)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
