package main

import (
	"backend_student/database"
	"backend_student/middleware"
	"backend_student/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup Database
	database.SetupDatabase()
	database.Migrate()

	// Setup Gin
	r := gin.Default()

	// Setup CORS Middleware
	r.Use(middleware.CORS())

	// Setup Routes
	routes.SetupRoutes(r)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server failed:", err)
	}
}
