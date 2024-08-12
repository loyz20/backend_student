package routes

import (
	"backend_student/config"
	"backend_student/controllers"
	"backend_student/middleware"
	"backend_student/repository"
	"backend_student/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Initialize repository, service, and controller
	studentRepo := repository.NewStudentRepository(config.DB)
	studentService := services.NewStudentService(studentRepo)
	studentController := controllers.NewStudentController(studentService)

	classRepo := repository.NewClassRepository(config.DB)
	classService := services.NewClassService(classRepo)
	classController := controllers.NewClassController(classService)

	userRepo := repository.NewUserRepository(config.DB)
	tokenRepo := repository.NewRefreshTokenRepository(config.DB)
	authService := services.NewAuthService(userRepo, studentRepo)
	tokenSevice := services.NewRefreshTokenService(tokenRepo)
	authController := controllers.NewAuthController(authService, tokenSevice)

	attendanceRepo := repository.NewAttendanceRepository(config.DB)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceController := controllers.NewAttendanceController(attendanceService)

	attendanceGroup := r.Group("/attendance")
	attendanceGroup.Use(middleware.AuthMiddleware())
	{
		attendanceGroup.POST("/", attendanceController.CreateAttendance)
		attendanceGroup.GET("/:id", attendanceController.GetAttendance)
		attendanceGroup.PUT("/:id", attendanceController.UpdateAttendance)
		attendanceGroup.DELETE("/:id", attendanceController.DeleteAttendance)
	}

	// Create a new route group for class
	classGroup := r.Group("/class")
	// classGroup.Use(middleware.AuthMiddleware()) // Apply your middleware here
	{
		classGroup.POST("/", classController.CreateClass)
		classGroup.GET("/:id", classController.GetClass)
		classGroup.PUT("/:id", classController.UpdateClass)
		classGroup.DELETE("/:id", classController.DeleteClass)
	}

	// Create a new route group for students
	studentGroup := r.Group("/students")
	studentGroup.Use(middleware.AuthMiddleware()) // Apply your middleware here
	{
		studentGroup.POST("/", studentController.CreateStudent)
		studentGroup.GET("/:id", studentController.GetStudent)
		studentGroup.PUT("/:id", studentController.UpdateStudent)
		studentGroup.DELETE("/:id", studentController.DeleteStudent)
	}

	// Create a new route group for authentication
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/verify-student", authController.VerifyStudent)
		authGroup.POST("/refresh-token", authController.RefreshToken)
	}
}
