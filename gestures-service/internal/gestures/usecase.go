package gestures

import (
	"context"

	"cyansnbrst/gestures-service/internal/models"
)

// Gestures usecase interface
type UseCase interface {
	GetGesture(ctx context.Context, id int64) (*models.Gesture, error)
	ListGestures(ctx context.Context) ([]models.Gesture, error)
	CreateGesture(ctx context.Context, gesture *models.Gesture) (int64, error)
	UpdateGesture(ctx context.Context, gesture *models.Gesture) error
	DeleteGesture(ctx context.Context, id int64) error
}
