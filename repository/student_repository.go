package repository

import (
	"backend_student/models"

	"gorm.io/gorm"
)

// StudentRepository defines the repository interface
type StudentRepository interface {
	Create(student *models.Student) error
	FindByID(id string) (*models.Student, error)
	Update(student *models.Student) error
	Delete(id string) error
}

// studentRepository implements StudentRepository interface
type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

// Create inserts a new Student into the database
func (r *studentRepository) Create(student *models.Student) error {
	result := r.db.Create(student)
	return result.Error
}

// FindByID retrieves a Student by ID
func (r *studentRepository) FindByID(id string) (*models.Student, error) {
	var student models.Student
	result := r.db.First(&student, "studentid = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

// Update updates an existing Student
func (r *studentRepository) Update(student *models.Student) error {
	result := r.db.Save(student)
	return result.Error
}

// Delete removes a Student from the database
func (r *studentRepository) Delete(id string) error {
	result := r.db.Delete(&models.Student{}, id)
	return result.Error
}
