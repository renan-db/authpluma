package routes

import (
	handler "project/internal/delivery/http/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	// Group user routes
	userGroup := e.Group("/users")
	
	// Register routes
	userGroup.POST("", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.GET("", userHandler.ListUsers)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)
} 