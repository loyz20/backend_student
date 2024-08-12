package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// ClassRepository defines the repository interface
type ClassRepository interface {
	Create(class *models.Class) error
	FindByID(id string) (*models.Class, error)
	Update(class *models.Class) error
	Delete(id string) error
}

// classRepository implements ClassRepository interface
type classRepository struct {
	db *gorm.DB
}

// NewClassRepository creates a new instance of ClassRepository
func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

// Create inserts a new Class into the database
func (r *classRepository) Create(class *models.Class) error {
	result := r.db.Create(class)
	return result.Error
}

// FindByID retrieves a Class by ID
func (r *classRepository) FindByID(id string) (*models.Class, error) {
	var class models.Class
	result := r.db.First(&class, "classid = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

// Update updates an existing Class
func (r *classRepository) Update(class *models.Class) error {
	result := r.db.Save(class)
	return result.Error
}

// Delete removes a Class from the database
func (r *classRepository) Delete(id string) error {
	result := r.db.Delete(&models.Class{}, id)
	return result.Error
}
