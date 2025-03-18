package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/preetheshkchordify/go-mongo-crud/controllers"
	"github.com/preetheshkchordify/go-mongo-crud/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add middleware
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.LoggingMiddleware())

	// Initialize controllers
	userController := controllers.NewUserController()

	// API routes
	api := router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/register", userController.RegisterUser)
			users.GET("/:id", userController.GetUserByID)
			users.GET("", userController.GetUserByEmail)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}

	return router
}
