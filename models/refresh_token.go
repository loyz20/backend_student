package models

import "gorm.io/gorm"

// RefreshToken defines the structure for the refresh token model
type RefreshToken struct {
	gorm.Model
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
