package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render a templ component
func renderView(c echo.Context, cmp templ.Component) error {
	// Set the content type to HMTL
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Writes the component to the echo context, and returns any errors
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
