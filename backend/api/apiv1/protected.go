package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func protected(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Protected"})
}
