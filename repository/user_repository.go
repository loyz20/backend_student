package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// UserRepository defines the structure for user repository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername retrieves a user by username
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser saves a new user to the database
func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}
