package handlers

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, ah *AuthHandler, th *TaskHandler) {
	e.GET("/", ah.flagsMiddleware(ah.homeHandler))
	e.GET("/login", ah.flagsMiddleware(ah.loginHandler))
	e.POST("/login", ah.flagsMiddleware(ah.loginHandler))
	e.GET("/register", ah.flagsMiddleware(ah.registerHandler))
	e.POST("/register", ah.flagsMiddleware(ah.registerHandler))

	// /* ↓ Protected Routes ↓ */
	e.POST("/logout", ah.authMiddleware(ah.logoutHandler))
	e.GET("/task", ah.authMiddleware(th.taskHandler))
	e.POST("/task", ah.authMiddleware(th.createTaskHandler))
	e.DELETE("/task/:id", ah.authMiddleware(th.deleteTaskHandler))
	e.GET("/task/table", ah.authMiddleware(th.taskTableHandler))
	e.GET("/task/:id/edit", ah.authMiddleware(th.taskTableHandler))

	e.PATCH("/task/:id", ah.authMiddleware(th.updateTaskHandler))
	e.PATCH("/task/:id/complete", ah.authMiddleware(th.completeTaskHandler))
	e.PATCH("/task/:id/review", ah.authMiddleware(th.reviewTaskHandler))

	// authGroup := e.Group("/auth", ah.authMiddleware)
	// authGroup.POST("/logout", ah.logoutHandler)
	//
	// protectedGroup := e.Group("/task", ah.authMiddleware)
	// protectedGroup.GET("/list", th.taskHandler)
	// protectedGroup.POST("/", th.createTaskHandler)
	// protectedGroup.PATCH("/:id", th.updateTaskHandler)
	// protectedGroup.PATCH("/:id/complete", th.completeTaskHandler)
	// protectedGroup.PATCH("/:id/review", th.reviewTaskHandler)
	// protectedGroup.DELETE("/:id", th.deleteTaskHandler)

	/* ↓ Fallback Page ↓ */
	e.GET("/*", RouteNotFoundHandler)
}
