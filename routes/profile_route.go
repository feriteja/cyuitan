// routes/user_routes.go
package routes

import (
	"github.com/feriteja/cyuitan/handlers"
	"github.com/feriteja/cyuitan/middleware"
	"github.com/gin-gonic/gin"
)

// Set up routes related to users
func SetupProfileRoutes(router *gin.Engine) {
	userRoutes := router.Group("/profile", middleware.AuthMiddleware())
	{
		userRoutes.PUT("/", handlers.EditProfile)
	}
}
