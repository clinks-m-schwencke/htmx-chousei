package handlers

import (
	"net/http"
	// "time"

	"github.com/labstack/echo/v4"
)

func handleMeetingPOST(c echo.Context) error {
	// title := c.FormValue("title")
	// description := c.FormValue("description")
	// dates := c.FormValue("dates")

	return c.Render(http.StatusOK, "meeting", nil)
}

func handleMeetingGET(c echo.Context) error {
	// TODO add data to template
	return c.Render(http.StatusOK, "meeting", nil)
}
