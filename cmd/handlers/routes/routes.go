package handlers

import (
	"chopitto-task/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.Static("/public", "public")

	// TODO
	// e.GET("/login", handlers.HandleIndexGet)

	e.GET("/", handlers.HandleIndexGet)

	// TODO
	// e.GET("task", handlers.HandleTaskGet) // Do I need this one?
	e.POST("task", handlers.HandleTaskPost)
	// e.PATCH("task", handlers.HandleTaskPatch)
	// e.DELETE("task", handlers.HandleTaskDelete)

	// e.POST("/meeting", handlers.HandleMeetingPost)

	// e.GET("/hello", func(c echo.Context) error {
	//
	// 	return c.Render(http.StatusOK, "index", nil)
	// })
	// e.GET("/meeting", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "meeting", nil)
	// })
	// e.POST("/meeting", func(c echo.Context) error {
	// 	// title := c.FormValue("title")
	// 	// datetimes := c.FormValue("datetimes")
	// 	// description := c.FormValue("description")
	// 	//
	// 	// if title == "" {
	// 	//
	// 	// }
	// 	//
	// 	// if page.Data.hasEmail(email) {
	// 	// 	formData := newFormData()
	// 	// 	formData.Values["name"] = name
	// 	// 	formData.Values["email"] = email
	// 	// 	formData.Errors["email"] = "Email already exists"
	// 	//
	// 	// 	return c.Render(http.StatusUnprocessableEntity, "form", formData)
	// 	// }
	// 	//
	// 	// contact := newContact(name, email)
	// 	// page.Data.Contacts = append(page.Data.Contacts, contact)
	// 	//
	// 	// c.Render(200, "form", newFormData())
	// 	// return c.Render(200, "oob-contact", contact)
	// 	// return c.Render(http.StatusOK, "meeting", person)
	// 	return nil
	// })
	//
	// e.GET("/meeting/new", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "meeting_new", nil)
	// })
}
