package services

import (
	"backend_student/models"
	"backend_student/repository"
)

// StudentService handles business logic for Students
type StudentService struct {
	repo repository.StudentRepository
}

// NewStudentService creates a new instance of StudentService
func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// CreateStudent creates a new Student
func (s *StudentService) CreateStudent(Student *models.Student) error {
	return s.repo.Create(Student)
}

// GetStudentByID retrieves a Student by ID
func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.repo.FindByID(id)
}

// UpdateStudent updates an existing Student
func (s *StudentService) UpdateStudent(Student *models.Student) error {
	return s.repo.Update(Student)
}

// DeleteStudent deletes a Student by ID
func (s *StudentService) DeleteStudent(id string) error {
	return s.repo.Delete(id)
}
