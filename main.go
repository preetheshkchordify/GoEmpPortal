package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/preetheshkchordify/go-mongo-crud/config"
	_ "github.com/preetheshkchordify/go-mongo-crud/docs" // Import generated Swagger docs
	"github.com/preetheshkchordify/go-mongo-crud/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Mongo CRUD API
// @version 1.0
// @description This is a sample CRUD API using Golang and MongoDB.
// @host localhost:8080
// @BasePath /api
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: Error loading .env file. Make sure it exists.")
	}

	// Initialize database connection
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer config.DisconnectDB()

	// Set up Gin router
	router := routes.SetupRouter()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server in a goroutine
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := router.Run(":" + port); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")
}
