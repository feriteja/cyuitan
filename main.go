// main.go
package main

import (
	"log"

	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/models"
	"github.com/feriteja/cyuitan/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database connection
	db, err := database.Init()
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate User model
	db.AutoMigrate(&models.User{}, &models.Auth{}, &models.Post{}, &models.Comment{}, &models.Profile{})

	// Initialize Gin router
	r := gin.Default()

	// Routes
	routes.SetupUserRoutes(r)

	// Start server
	r.Run(":8080")
}
