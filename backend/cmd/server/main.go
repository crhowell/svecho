package main

import (
	"backend/api/apiv1"
	"backend/api/authv1"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// CORS Allowed Origins
	origins := make([]string, 0)
	allowed := os.Getenv("ALLOWED_ORIGINS")
	for _, origin := range strings.Split(allowed, ",") {
		origins = append(origins, strings.TrimSpace(origin))
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
	}))

	e.Use(middleware.Logger())

	// register auth v1 routes from authv1 package
	authv1Group := e.Group("/api/auth/v1")
	authv1.Register(authv1Group)

	// register api v1 routes from apiv1 package
	apiv1Group := e.Group("/api/v1")
	apiv1.Register(apiv1Group)

	e.Logger.Fatal(e.Start(":8000"))
}
