package controllers

import (
	"net/http"
	"strconv"

	"backend_student/services"

	"github.com/gin-gonic/gin"
)

func CreateAttendance(c *gin.Context) {
	username := c.PostForm("username")
	latitude, _ := strconv.ParseFloat(c.PostForm("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.PostForm("longitude"), 64)
	photoURL := c.PostForm("photo_url")

	attendance, err := services.CreateAttendance(username, latitude, longitude, photoURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": attendance})
}

func GetMonthlyReport(c *gin.Context) {
	username := c.Param("username")
	report, err := services.GetMonthlyReport(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": report})
}
