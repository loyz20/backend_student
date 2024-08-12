package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

// userRepository struct implementing UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// FindByUsername retrieves a user by username
func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser saves a new user to the database
func (repo *userRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}
