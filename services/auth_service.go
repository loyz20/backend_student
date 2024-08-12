package services

import (
	"backend_student/models"
	"backend_student/repository"

	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the structure for authentication service
type AuthService struct {
	userRepo *repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// RegisterUser registers a new user with hashed password
func (service *AuthService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{Username: username, Password: string(hashedPassword)}
	return service.userRepo.CreateUser(user)
}

// AuthenticateUser authenticates a user by checking password
func (service *AuthService) AuthenticateUser(username, password string) (bool, error) {
	user, err := service.userRepo.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, nil
	}
	return true, nil
}
