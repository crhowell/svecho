package apiv1

import (
	"backend/api/middleware"

	"github.com/labstack/echo/v4"
)

// Register registers all /api/v1 routes and applies middleware (if any)
func Register(e *echo.Group) {
	e.GET("/status", status)

	// Protected routes using JWT claims
	protectedGroup := e.Group("", middleware.JWTMiddleware())
	protectedGroup.GET("/protected", protected)
}
