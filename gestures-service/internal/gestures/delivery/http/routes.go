package http

import (
	"cyansnbrst/gestures-service/internal/gestures"

	"github.com/labstack/echo/v4"
)

// Register gestures routes
func RegisterGesturesRoutes(g *echo.Group, h gestures.Handlers) {
	g.GET("/", h.ListGestures)
	g.GET("/:id", h.GetGesture)
	g.POST("/", h.CreateGesture)
	g.DELETE("/:id", h.DeleteGesture)
}
