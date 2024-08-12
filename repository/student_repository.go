package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// StudentRepository defines the repository interface
type StudentRepository interface {
	Create(Student *models.Student) error
	FindByID(id string) (*models.Student, error)
	Update(Student *models.Student) error
	Delete(id string) error
}

// StudentRepository implements StudentRepository interface
type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

// Create inserts a new Student into the database
func (r *studentRepository) Create(Student *models.Student) error {
	result := r.db.Create(Student)
	return result.Error
}

// FindByID retrieves a Student by ID
func (r *studentRepository) FindByID(id string) (*models.Student, error) {
	var Student models.Student
	result := r.db.First(&Student, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Student, nil
}

// Update updates an existing Student
func (r *studentRepository) Update(Student *models.Student) error {
	result := r.db.Save(Student)
	return result.Error
}

// Delete removes a Student from the database
func (r *studentRepository) Delete(id string) error {
	result := r.db.Delete(&models.Student{}, id)
	return result.Error
}
