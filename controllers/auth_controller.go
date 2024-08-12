package controllers

import (
	"backend_student/services"
	"backend_student/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController defines the structure for authentication controller
type AuthController struct {
	authService         *services.AuthService
	refreshTokenService *services.RefreshTokenService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register handles user registration
func (ctrl *AuthController) Register(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if err := ctrl.authService.RegisterUser(request.Username, request.Password); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "User registered successfully", nil)
}

// Login handles user login and returns a JWT token
func (ctrl *AuthController) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	isAuthenticated, err := ctrl.authService.AuthenticateUser(request.Username, request.Password)
	if err != nil || !isAuthenticated {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	token, err := utils.GenerateToken(request.Username)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Login successful", gin.H{"token": token})
}

// RefreshToken handles refresh token requests
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	userID, err := ctrl.refreshTokenService.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}

	// Generate new access token
	newToken, err := utils.GenerateToken(userID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "success", "Token refreshed successfully", gin.H{"token": newToken})
}
