package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"htmx-chousei.com/name/cmd/handlers"
)

func Routes(e *echo.Echo) {

	e.Static("/public", "public")

	e.GET("/", handlers.HandleIndex)

	e.GET("/hello", func(c echo.Context) error {

		return c.Render(http.StatusOK, "index", nil)
	})
	e.GET("/meeting", func(c echo.Context) error {
		return c.Render(http.StatusOK, "meeting", nil)
	})
	e.POST("/meeting", func(c echo.Context) error {
		// title := c.FormValue("title")
		// datetimes := c.FormValue("datetimes")
		// description := c.FormValue("description")
		//
		// if title == "" {
		//
		// }
		//
		// if page.Data.hasEmail(email) {
		// 	formData := newFormData()
		// 	formData.Values["name"] = name
		// 	formData.Values["email"] = email
		// 	formData.Errors["email"] = "Email already exists"
		//
		// 	return c.Render(http.StatusUnprocessableEntity, "form", formData)
		// }
		//
		// contact := newContact(name, email)
		// page.Data.Contacts = append(page.Data.Contacts, contact)
		//
		// c.Render(200, "form", newFormData())
		// return c.Render(200, "oob-contact", contact)
		// return c.Render(http.StatusOK, "meeting", person)
		return nil
	})

	e.GET("/meeting/new", func(c echo.Context) error {
		return c.Render(http.StatusOK, "meeting_new", nil)
	})
}
