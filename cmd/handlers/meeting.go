package handlers

import (
	"net/http"
	"strings"
	// "time"

	"chopitto-task/cmd/lib/types"
	"chopitto-task/views"
	"github.com/labstack/echo/v4"
)

func HandleMeetingPost(c echo.Context) error {
	title := strings.TrimSpace(c.FormValue("title"))
	description := strings.TrimSpace(c.FormValue("description"))
	dates := strings.TrimSpace(c.FormValue("dates"))

	if title == "" || dates == "" {
		formData := types.NewFormData()
		formData.Values["title"] = title
		formData.Values["description"] = description
		formData.Values["dates"] = dates
		return Render(c, http.StatusUnprocessableEntity, views.MeetingForm(formData))
	}
	Render(c, http.StatusOK, views.OobMeetingCard(title, description))
	return Render(c, http.StatusOK, views.MeetingForm(types.NewFormData()))
}

func HandleMeetingGET(c echo.Context) error {
	// TODO add data to template
	return c.Render(http.StatusOK, "meeting", nil)
}
