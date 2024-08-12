package services

import (
	"backend_student/models"
	"backend_student/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the structure for authentication service
type AuthService struct {
	userRepo    repository.UserRepository
	studentRepo repository.StudentRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo repository.UserRepository, studentRepo repository.StudentRepository) *AuthService {
	return &AuthService{userRepo: userRepo, studentRepo: studentRepo}
}

// RegisterUser registers a new user with hashed password
func (service *AuthService) RegisterUser(username, password string) error {
	// Check if user already exists
	existingUser, _ := service.userRepo.FindByUsername(username)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user
	user := &models.User{Username: username, Password: string(hashedPassword)}
	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}

	// Find student by username
	student, err := service.studentRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	if student == nil {
		return errors.New("student not found")
	}

	// Update student's username
	student.Username = username
	if err := service.studentRepo.Update(student); err != nil {
		return err
	}

	return nil
}

// AuthenticateUser authenticates a user by checking password
func (service *AuthService) AuthenticateUser(username, password string) (bool, error) {
	// Find user by username
	user, err := service.userRepo.FindByUsername(username)
	if err != nil {
		return false, err
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

// GetStudentByID retrieves a Student by ID
func (service *AuthService) GetStudentByID(id string) (*models.Student, error) {
	// Find student by ID
	student, err := service.studentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}
