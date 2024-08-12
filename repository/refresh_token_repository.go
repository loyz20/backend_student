package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// RefreshTokenRepository defines the structure for refresh token repository
type RefreshTokenRepository struct {
	db *gorm.DB
}

// NewRefreshTokenRepository creates a new instance of RefreshTokenRepository
func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// CreateToken saves a new refresh token to the database
func (repo *RefreshTokenRepository) CreateToken(refreshToken *models.RefreshToken) error {
	return repo.db.Create(refreshToken).Error
}

// FindByToken retrieves a refresh token by its value
func (repo *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	if err := repo.db.Where("token = ?", token).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

// DeleteToken removes a refresh token from the database
func (repo *RefreshTokenRepository) DeleteToken(token string) error {
	return repo.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}
