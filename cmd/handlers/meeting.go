package handlers

import (
	"net/http"
	// "time"

	"github.com/labstack/echo/v4"
)

type Meeting struct {
	Title       string
	Description string
	// DateTime *time.Time.list
	Dates string
}

func newMeeting(title string, description string, dates string) Meeting {
	return Meeting{
		Title:       title,
		Description: description,
		Dates:       dates,
	}
}

func handleMeetingNew(c echo.Context) error {
	return c.Render(http.StatusOK, "meeting", nil)
}

func handleMeetingGET(c echo.Context) error {
	// TODO add data to template
	return c.Render(http.StatusOK, "meeting", nil)
}
