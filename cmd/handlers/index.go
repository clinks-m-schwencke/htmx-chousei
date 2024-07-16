package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"htmx-chousei.com/name/cmd/stores"
)

func HandleIndex(c echo.Context) error {
	store := stores.SessionDataStore
	return c.Render(http.StatusOK, "index", store)
}
