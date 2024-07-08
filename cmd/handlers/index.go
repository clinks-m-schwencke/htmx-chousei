package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
