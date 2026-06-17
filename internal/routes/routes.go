package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/indahcs/go-gin-api-boilerplate/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")

	{
		api.GET("/health", handlers.HealthCheck)
	}
}
