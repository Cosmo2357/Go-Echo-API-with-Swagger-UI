package main

import (
	"myapp/config"
	_ "myapp/docs" // Import generated docs
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title My API Example
// @version 1.0
// @description This is a sample server for demonstrating Go Echo with Swagger.
// @host localhost:8080
// @BasePath /
func main() {
	config.LoadConfig() // Load configuration
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.InitRoutes(e) // Initialize routes

	e.Logger.Fatal(e.Start(":8080"))
}
