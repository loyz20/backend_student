package controllers

import (
	"backend_student/models"
	"backend_student/services"
	"backend_student/utils" // Import package utils
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StudentController defines the structure for student-related endpoints
type StudentController struct {
	studentService *services.StudentService
}

// NewStudentController creates a new instance of StudentController
func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{studentService: studentService}
}

// CreateStudent handles the creation of a new student
func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := ctrl.studentService.CreateStudent(&student); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create student", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Student created successfully", nil)
}

// GetStudent retrieves a student by ID
func (ctrl *StudentController) GetStudent(c *gin.Context) {
	id := c.Param("id")
	student, err := ctrl.studentService.GetStudentByID(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Student not found", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Student retrieved successfully", student)
}

// UpdateStudent handles updating an existing student
func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")

	// Convert string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}
	student.ID = uint(id)

	if err := ctrl.studentService.UpdateStudent(&student); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update student", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Student updated successfully", nil)
}

// DeleteStudent handles deleting a student by ID
func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.studentService.DeleteStudent(id); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete student", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Student deleted successfully", nil)
}
