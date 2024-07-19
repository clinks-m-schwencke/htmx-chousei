package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// This custom render replaces Echo's echo.Context.Render() with templ's templ.Component.Render()
func Render(ctx echo.Context, statusCode int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := component.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
