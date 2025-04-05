package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	gesturesHTTP "cyansnbrst/gestures-service/internal/gestures/delivery/http"
	gesturesRepository "cyansnbrst/gestures-service/internal/gestures/repository"
	gesturesUseCase "cyansnbrst/gestures-service/internal/gestures/usecase"
	"cyansnbrst/gestures-service/pkg/validator"
)

// Register server handlers
func (s *Server) RegisterHandlers() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.Validator = validator.NewCustomValidator()

	gesturesRepo := gesturesRepository.NewGesturesRepo(s.db)

	gesturesUC := gesturesUseCase.NewGesturesUseCase(gesturesRepo)

	gesturesHandlers := gesturesHTTP.NewGesturesHandlers(gesturesUC, s.logger)

	api := e.Group("/api")

	gesturesHTTP.RegisterGesturesRoutes(api, gesturesHandlers)

	return e
}
