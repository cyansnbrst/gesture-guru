package usecase

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
	"github.com/cyansnbrst/gesture-guru/gestures-service/pkg/db"
	"github.com/cyansnbrst/gesture-guru/gestures-service/pkg/logger"
)

// Gestures usecase struct
type gesturesUseCase struct {
	logger       *zap.Logger
	gesturesRepo gestures.Repository
}

var (
	ErrGesturesNotFound = errors.New("no gestures in requested category")
)

// NewGesturesUseCase - конструктор для слоя юзкейсов
func NewGesturesUseCase(logger *zap.Logger, gesturesRepo gestures.Repository) gestures.UseCase {
	return &gesturesUseCase{logger: logger, gesturesRepo: gesturesRepo}
}

func (uc *gesturesUseCase) GetGesture(ctx context.Context, id int64) (*models.Gesture, error) {
	const op = "Gestures.GetGesture"

	gesture, err := uc.gesturesRepo.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, db.ErrGestureNotFound) {
			logger.LogError(uc.logger, op, err, "failed to get gesture by id")
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return gesture, nil
}

func (uc *gesturesUseCase) ListGestures(ctx context.Context) ([]models.Gesture, error) {
	const op = "Gestures.ListGestures"

	gestures, err := uc.gesturesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(gestures) == 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrGesturesNotFound)
	}

	return gestures, nil
}

func (uc *gesturesUseCase) CreateGesture(ctx context.Context, gesture *models.Gesture) (int64, error) {
	const op = "Gestures.CreateGesture"

	id, err := uc.gesturesRepo.Create(ctx, gesture)
	if err != nil {
		if !errors.Is(err, db.ErrInvalidCategory) {
			logger.LogError(uc.logger, op, err, "failed to create gesture")
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (uc *gesturesUseCase) UpdateGesture(ctx context.Context, gesture *models.Gesture) error {
	const op = "Gestures.UpdateGesture"

	err := uc.gesturesRepo.Update(ctx, gesture)
	if err != nil {
		if !errors.Is(err, db.ErrInvalidCategory) && !errors.Is(err, db.ErrGestureNotFound) {
			logger.LogError(uc.logger, op, err, "failed to update gesture")
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (uc *gesturesUseCase) DeleteGesture(ctx context.Context, id int64) error {
	const op = "Gestures.DeleteGesture"

	err := uc.gesturesRepo.Delete(ctx, id)
	if err != nil {
		if !errors.Is(err, db.ErrGestureNotFound) {
			logger.LogError(uc.logger, op, err, "failed to delete gesture")
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
