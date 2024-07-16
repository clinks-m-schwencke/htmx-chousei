package main

import (
	// "context"
	"io"
	// "os"
	"text/template"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	// "htmx-chousei.com/name/cmd/routes"
	"htmx-chousei.com/name/views"
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

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	// component := views.Hello("John")
	// component.Render(context.Background(), os.Stdout)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
		return render(c, views.Hello("John"))
	})
	//
	// NewTemplateRenderer(e, "views/*.html")
	//
	// routes.Routes(e)
	//
	e.Logger.Fatal(e.Start(":8889"))
}
