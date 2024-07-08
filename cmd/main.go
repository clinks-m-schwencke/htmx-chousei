package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"htmx-chousei.com/name/cmd/routes"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(writer io.Writer, name string, data interface{}, context echo.Context) error {
	return t.Templates.ExecuteTemplate(writer, name, data)
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template{
		Templates: templates,
	}
}

func NewTemplateRenderer(e *echo.Echo, paths ...string) {
	templates := &template.Template{}
	for i := range paths {
		template.Must(templates.ParseGlob(paths[i]))
	}
	t := newTemplate(templates)
	e.Renderer = t
}

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	NewTemplateRenderer(e, "views/*.html")

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":8889"))
}
