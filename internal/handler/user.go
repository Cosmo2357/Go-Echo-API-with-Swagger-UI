package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser handles GET requests to retrieve user details
// @Summary Get user information
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":   id,
		"name": "John Doe",
	})
}
