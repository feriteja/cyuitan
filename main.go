// main.go
package main

import (
	"github.com/feriteja/cyuitan/database"
	"github.com/feriteja/cyuitan/handlers"
	"github.com/feriteja/cyuitan/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db, err := database.Init()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// AutoMigrate User model
	db.AutoMigrate(&models.User{})

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// Start server
	r.Run(":8080")
}
