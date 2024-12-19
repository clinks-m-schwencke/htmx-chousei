package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"chopitto-task/db"
	"chopitto-task/handlers"
	"chopitto-task/services"
)

// TODO: These should go into an .env file
const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "task.db"
)

func main() {

	e := echo.New()

	// Create static routes
	e.Static("/static", "static")

	// Setup custom error handling
	e.HTTPErrorHandler = handlers.CustomHttpErrorHandler

	// Verbose logging middleware
	e.Pre(middleware.Logger())

	// Centralised error handling middleware
	e.Use(middleware.Recover())

	// Remove final trailing slash in URL middleware
	e.Pre(middleware.RemoveTrailingSlash())

	// Rate limiter middleware - No SPAM please!
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	// Session Middleware
	e.Use(session.Middleware(sessions.NewCookieStore(([]byte(SECRET_KEY)))))

	// Create database store
	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatal("failed to create store: %s", err)
	}

	// Setup db interaction services and route handlers
	personService := services.NewPersonServices(services.Person{}, store)
	authHandler := handlers.NewAuthHandler(personService)
	//
	taskService := services.NewTaskServices(services.Task{}, store)
	taskHandler := handlers.NewTaskHandler(taskService)
	//
	handlers.SetupRoutes(e, authHandler, taskHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
