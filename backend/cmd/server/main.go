package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &map[string]string{"message": "Hello from the Echo backend!"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
