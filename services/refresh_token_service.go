package services

import (
	"backend_student/models"
	"backend_student/repository"
	"backend_student/utils"
)

// RefreshTokenService defines the structure for refresh token service
type RefreshTokenService struct {
	refreshTokenRepo *repository.RefreshTokenRepository
}

// NewRefreshTokenService creates a new instance of RefreshTokenService
func NewRefreshTokenService(refreshTokenRepo *repository.RefreshTokenRepository) *RefreshTokenService {
	return &RefreshTokenService{refreshTokenRepo: refreshTokenRepo}
}

// CreateRefreshToken creates and stores a new refresh token
func (service *RefreshTokenService) CreateRefreshToken(userID string) (string, error) {
	token, err := utils.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	rt := &models.RefreshToken{
		UserID: userID,
		Token:  token,
	}

	if err := service.refreshTokenRepo.CreateToken(rt); err != nil {
		return "", err
	}

	return token, nil
}

// ValidateRefreshToken checks if a refresh token is valid
func (service *RefreshTokenService) ValidateRefreshToken(token string) (string, error) {
	rt, err := service.refreshTokenRepo.FindByToken(token)
	if err != nil {
		return "", err
	}

	// Optionally, you can also validate if the token is not expired
	// if expired(rt) {
	//     return "", errors.New("token expired")
	// }

	return rt.UserID, nil
}

// DeleteRefreshToken removes a refresh token
func (service *RefreshTokenService) DeleteRefreshToken(token string) error {
	return service.refreshTokenRepo.DeleteToken(token)
}
