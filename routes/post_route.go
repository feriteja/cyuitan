// routes/user_routes.go
package routes

import (
	"github.com/feriteja/cyuitan/handlers"
	"github.com/feriteja/cyuitan/middleware"
	"github.com/gin-gonic/gin"
)

// Set up routes related to users
func SetupPostRoutes(router *gin.Engine) {
	postRoute := router.Group("/post", middleware.AuthMiddleware())
	{
		postRoute.POST("/", handlers.SendPost)
		postRoute.POST("/edit", handlers.EditPost)
	}
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("/", handlers.GetPosts)
	}
}
