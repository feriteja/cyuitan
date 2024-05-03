// routes/user_routes.go
package routes

import (
	"github.com/feriteja/cyuitan/handlers"
	"github.com/feriteja/cyuitan/middleware"
	"github.com/gin-gonic/gin"
)

// Set up routes related to users
func SetupPostRoutes(router *gin.Engine) {
	userRoutes := router.Group("/post", middleware.AuthMiddleware())
	{
		userRoutes.POST("/", handlers.SendPost)
	}
}
