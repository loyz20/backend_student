// controllers/user_controller.go
package controllers

import (
	"backend_student/database"
	"backend_student/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.User

	type SuccessResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Code:    "INVALID_INPUT",
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	var user models.User
	user.Username = input.Username
	if err := user.SetPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    "PASSWORD_HASHING_FAILED",
			Message: "Could not hash password",
		})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    "USER_CREATION_FAILED",
			Message: "Could not create user",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Status:  "success",
		Message: "User created successfully",
	})
}
