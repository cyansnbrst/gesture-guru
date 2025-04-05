package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"cyansnbrst/gestures-service/internal/gestures"
	"cyansnbrst/gestures-service/internal/models"
	hh "cyansnbrst/gestures-service/pkg/http_helpers"
)

// Gestures handlers struct
type gesturesHandlers struct {
	useCase gestures.UseCase
	logger  *zap.Logger
}

// NewGesturesHandlers - конструктор для хендлеров жестов
func NewGesturesHandlers(useCase gestures.UseCase, logger *zap.Logger) gestures.Handlers {
	return &gesturesHandlers{
		useCase: useCase,
		logger:  logger,
	}
}

// GetGesture - обработчик для получения жеста по ID
func (h *gesturesHandlers) GetGesture(c echo.Context) error {
	id, err := hh.ParseIDParam(c, "id")
	if err != nil {
		return hh.BadRequestResponse(c, err)
	}

	gesture, err := h.useCase.GetGesture(c.Request().Context(), id)
	if err != nil {
		return hh.ServerErrorResponse(c, h.logger, err)
	}

	return c.JSON(http.StatusOK, gesture)
}

// ListGestures - обработчик для получения списка жестов
func (h *gesturesHandlers) ListGestures(c echo.Context) error {
	gestures, err := h.useCase.ListGestures(c.Request().Context())
	if err != nil {
		return hh.ServerErrorResponse(c, h.logger, err)
	}

	return c.JSON(http.StatusOK, gestures)
}

// CreateGesture - обработчик для создания нового жеста
func (h *gesturesHandlers) CreateGesture(c echo.Context) error {
	var gesture models.Gesture
	if err := c.Bind(&gesture); err != nil {
		return hh.BadRequestResponse(c, err)
	}

	id, err := h.useCase.CreateGesture(c.Request().Context(), &gesture)
	if err != nil {
		return hh.ServerErrorResponse(c, h.logger, err)
	}

	return c.JSON(http.StatusCreated, map[string]int64{"id": id})
}

// // UpdateGesture - обработчик для обновления существующего жеста
// func (h *gesturesHandlers) UpdateGesture(c echo.Context) error {
// 	var gesture models.Gesture
// 	if err := c.Bind(&gesture); err != nil {
// 		return hh.BadRequestResponse(c, err)
// 	}

// 	if err := h.useCase.UpdateGesture(c.Request().Context(), &gesture); err != nil {
// 		return hh.ServerErrorResponse(c, h.logger, err)
// 	}

// 	return c.NoContent(http.StatusOK)
// }

// DeleteGesture - обработчик для удаления жеста по ID
func (h *gesturesHandlers) DeleteGesture(c echo.Context) error {
	id, err := hh.ParseIDParam(c, "id")
	if err != nil {
		return hh.BadRequestResponse(c, err)
	}

	if err := h.useCase.DeleteGesture(c.Request().Context(), id); err != nil {
		return hh.ServerErrorResponse(c, h.logger, err)
	}

	return c.NoContent(http.StatusOK)
}
