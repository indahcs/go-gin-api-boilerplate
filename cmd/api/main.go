package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indahcs/go-gin-api-boilerplate/internal/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
