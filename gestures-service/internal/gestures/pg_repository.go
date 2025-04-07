package gestures

import (
	"context"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
)

// Gestures repository interface
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Gesture, error)
	GetAll(ctx context.Context) ([]models.Gesture, error)
	Create(ctx context.Context, gesture *models.Gesture) (int64, error)
	Update(ctx context.Context, gesture *models.Gesture) error
	Delete(ctx context.Context, id int64) error
}
