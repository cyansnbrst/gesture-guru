package gestures

import "github.com/labstack/echo/v4"

// Merch handlers interface
type Handlers interface {
	GetGesture(c echo.Context) error
	ListGestures(c echo.Context) error
	CreateGesture(c echo.Context) error
	// UpdateGesture(c echo.Context) error
	DeleteGesture(c echo.Context) error
}
