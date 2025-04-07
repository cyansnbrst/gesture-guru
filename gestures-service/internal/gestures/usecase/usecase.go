package usecase

import (
	"context"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
)

// Gestures usecase struct
type gesturesUseCase struct {
	gesturesRepo gestures.Repository
}

// NewGesturesUseCase - конструктор для слоя юзкейсов
func NewGesturesUseCase(gesturesRepo gestures.Repository) gestures.UseCase {
	return &gesturesUseCase{gesturesRepo: gesturesRepo}
}

func (uc *gesturesUseCase) GetGesture(ctx context.Context, id int64) (*models.Gesture, error) {
	return uc.gesturesRepo.GetByID(ctx, id)
}

func (uc *gesturesUseCase) ListGestures(ctx context.Context) ([]models.Gesture, error) {
	return uc.gesturesRepo.GetAll(ctx)
}

func (uc *gesturesUseCase) CreateGesture(ctx context.Context, gesture *models.Gesture) (int64, error) {
	return uc.gesturesRepo.Create(ctx, gesture)
}

func (uc *gesturesUseCase) UpdateGesture(ctx context.Context, gesture *models.Gesture) error {
	return uc.gesturesRepo.Update(ctx, gesture)
}

func (uc *gesturesUseCase) DeleteGesture(ctx context.Context, id int64) error {
	return uc.gesturesRepo.Delete(ctx, id)
}
