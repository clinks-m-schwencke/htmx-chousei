package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"htmx-chousei.com/name/views"
)

func HandleIndexGet(c echo.Context) error {
	return Render(c, http.StatusOK, views.Index())
}
