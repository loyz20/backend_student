package controllers

import (
	"backend_student/models"
	"backend_student/services"
	"backend_student/utils" // Import package utils
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ClassController defines the structure for class-related endpoints
type ClassController struct {
	classService *services.ClassService
}

// NewClassController creates a new instance of ClassController
func NewClassController(classService *services.ClassService) *ClassController {
	return &ClassController{classService: classService}
}

// CreateClass handles the creation of a new class
func (ctrl *ClassController) CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := ctrl.classService.CreateClass(&class); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create class", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Class created successfully", nil)
}

// GetClass retrieves a class by ID
func (ctrl *ClassController) GetClass(c *gin.Context) {
	id := c.Param("id")
	class, err := ctrl.classService.GetClassByID(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Class not found", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Class retrieved successfully", class)
}

// UpdateClass handles updating an existing class
func (ctrl *ClassController) UpdateClass(c *gin.Context) {
	idStr := c.Param("id")

	// Convert string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}
	class.ID = uint(id)

	if err := ctrl.classService.UpdateClass(&class); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update class", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Class updated successfully", nil)
}

// DeleteClass handles deleting a class by ID
func (ctrl *ClassController) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.classService.DeleteClass(id); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete class", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Class deleted successfully", nil)
}
