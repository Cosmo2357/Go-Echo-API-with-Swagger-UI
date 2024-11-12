package routes

import (
	"myapp/internal/handler"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes the routes for the application
func InitRoutes(e *echo.Echo) {
	e.GET("/users/:id", handler.GetUser)
}
