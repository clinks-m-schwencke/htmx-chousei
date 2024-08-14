package handlers

import (
	"net/http"
	"strings"

	"chopitto-task/cmd/lib/types"
	"chopitto-task/views"
	"github.com/labstack/echo/v4"
)

func HandleSignUpPost(c echo.Context) error {
	name := strings.TrimSpace(c.FormValue("name"))
	email := strings.TrimSpace(c.FormValue("email"))
	password := strings.TrimSpace(c.FormValue("password"))

	if email == "" || password == "" || name == "" {
		formData := types.NewFormData()
		formData.Values["name"] = name
		formData.Values["email"] = email
		formData.Values["password"] = password
		// TODO: Change MeetingForm to sign up
		return Render(c, http.StatusUnprocessableEntity, views.MeetingForm(formData))
	}
	// TODO: Change MeetingForm to sign up
	return Render(c, http.StatusOK, views.MeetingForm(types.NewFormData()))
}
