package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"chopitto-task/cmd/handlers/routes"
	data "chopitto-task/cmd/lib/data"
)

func main() {

	e := echo.New()

	data.OpenDatabase()

	// db, err := sql.Open("sqlite3", "./tasks.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	handlers.Routes(e)

	e.Logger.Fatal(e.Start(":8889"))
}
