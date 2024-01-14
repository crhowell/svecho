package authv1

import (
	"backend/api/middleware"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	var creds loginData
	if err := c.Bind(&creds); err != nil {
		return err
	}

	// TODO: Logic to check username/password from database/service
	if creds.Username != "svecho" || creds.Password != "supersecret" {
		return echo.ErrUnauthorized
	}

	// Set custom claims and sign token for client to use for authentication
	claims := &middleware.JWTCustomClaims{
		Name:  "Svelte Echo User",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET_KEY")
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
