package controllers

import (
	"backend_student/database"
	"backend_student/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "your_secret_key" // Ganti dengan secret key yang lebih aman

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Code:    "INVALID_INPUT",
			Message: "Invalid input: " + err.Error(),
		})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Status:  "error",
			Code:    "INVALID_CREDENTIALS",
			Message: "Invalid username or password",
		})
		return
	}

	if !user.CheckPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Status:  "error",
			Code:    "INVALID_CREDENTIALS",
			Message: "Invalid username or password",
		})
		return
	}

	token, err := generateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Code:    "TOKEN_GENERATION_FAILED",
			Message: "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
