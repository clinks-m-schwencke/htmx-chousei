package handlers

import (
	"net/http"

	"chopitto-task/views"
	"github.com/labstack/echo/v4"
)

func HandleHelloGET(c echo.Context) error {
	return Render(c, http.StatusOK, views.Who())
}

func HandleHelloPOST(c echo.Context) error {
	name := c.FormValue("name")
	return Render(c, http.StatusOK, views.Hello(name))
}
