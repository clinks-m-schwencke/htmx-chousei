package main

import (
	// "context"
	// "io"
	// "os"
	// "text/template"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"htmx-chousei.com/name/cmd/handlers/routes"
)

func main() {

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	handlers.Routes(e)

	e.Logger.Fatal(e.Start(":8889"))
}
