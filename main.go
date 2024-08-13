package main

import (
	"backend_student/config"
	"backend_student/middleware"
	"backend_student/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	config.InitializeConfig()

	// Create a new Gin router
	r := gin.Default()

	// Setup CORS
	r.Use(middleware.CORS())

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
