// routes/user_routes.go
package routes

import (
	"github.com/feriteja/cyuitan/handlers"
	"github.com/gin-gonic/gin"
)

// Set up routes related to authentication
func SetupAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
		// Add other auth routes as needed
	}
}
