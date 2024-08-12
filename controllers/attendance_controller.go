package controllers

import (
	"backend_student/models"
	"backend_student/services"
	"backend_student/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceController struct {
	attendanceService services.AttendanceService
}

func NewAttendanceController(attendanceService services.AttendanceService) *AttendanceController {
	return &AttendanceController{attendanceService}
}

func (ctrl *AttendanceController) CreateAttendance(c *gin.Context) {
	var attendance models.Attendance

	// Bind JSON for attendance details
	if err := c.ShouldBindJSON(&attendance); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Handle file upload
	file, err := c.FormFile("selfie_image")
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to upload image", err)
		return
	}

	// Generate a unique filename
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), attendance.StudentID, filepath.Ext(file.Filename))
	savePath := filepath.Join("uploads", "selfies", filename)

	// Save the file to disk
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to save image", err)
		return
	}

	// Update attendance struct with the file URL or path
	attendance.SelfieImageUrl = savePath

	// Calculate late minutes (assuming 8:00 AM is the threshold for being late)
	thresholdTime := time.Date(attendance.AttendanceDate.Year(), attendance.AttendanceDate.Month(), attendance.AttendanceDate.Day(), 8, 0, 0, 0, attendance.AttendanceDate.Location())
	if attendance.AttendanceDate.After(thresholdTime) && attendance.AttendanceType == "masuk" {
		duration := attendance.AttendanceDate.Sub(thresholdTime)
		attendance.LateMinutes = int(duration.Minutes())
	}

	// Save attendance record
	if err := ctrl.attendanceService.CreateAttendance(&attendance); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create attendance", err)
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Attendance created successfully", "", attendance)
}

func (c *AttendanceController) GenerateAttendanceReport(ctx *gin.Context) {
	studentID := ctx.Query("student_id")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid start date format", err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "Invalid end date format", err)
		return
	}

	report, err := c.attendanceService.GenerateAttendanceReport(studentID, startDate, endDate)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, "Failed to generate report", err)
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, "Attendance report generated successfully", "", report)
}

func (ctrl *AttendanceController) GetAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	attendance, err := ctrl.attendanceService.GetAttendanceByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Attendance not found", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Attendance retrieved successfully", "", attendance)
}

func (ctrl *AttendanceController) UpdateAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	var attendance models.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	attendance.ID = uint(id)
	if err := ctrl.attendanceService.UpdateAttendance(&attendance); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update attendance", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Attendance updated successfully", "", attendance)
}

func (ctrl *AttendanceController) DeleteAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	if err := ctrl.attendanceService.DeleteAttendance(uint(id)); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete attendance", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Attendance deleted successfully", "", nil)
}
