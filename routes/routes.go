package routes

import (
	"go-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.GET("/login", handlers.Login)
	}
}
