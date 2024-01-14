package authv1

import "github.com/labstack/echo/v4"

// Register registers all /api/auth/v1 routes and applies middleware (if any)
func Register(e *echo.Group) {
	e.POST("/login", login)
}
