package routes

import (
	"backend_student/controllers"
	"backend_student/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/users", controllers.CreateUser)

		// Middleware untuk rute yang memerlukan autentikasi
		auth := api.Group("/auth")
		auth.Use(middleware.AuthMiddleware())
		{
			// Tambahkan rute yang memerlukan autentikasi di sini
			// auth.GET("/profile", controllers.Profile)
			auth.POST("/attendance", controllers.CreateAttendance)
			auth.GET("/attendance/report/:username", controllers.GetMonthlyReport)
		}
	}
}
