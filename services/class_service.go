package services

import (
	"backend_student/models"
	"backend_student/repository"
)

// ClassService handles business logic for Classs
type ClassService struct {
	repo repository.ClassRepository
}

// NewClassService creates a new instance of ClassService
func NewClassService(repo repository.ClassRepository) *ClassService {
	return &ClassService{repo: repo}
}

// CreateClass creates a new Class
func (s *ClassService) CreateClass(Class *models.Class) error {
	return s.repo.Create(Class)
}

// GetClassByID retrieves a Class by ID
func (s *ClassService) GetClassByID(id string) (*models.Class, error) {
	return s.repo.FindByID(id)
}

// UpdateClass updates an existing Class
func (s *ClassService) UpdateClass(Class *models.Class) error {
	return s.repo.Update(Class)
}

// DeleteClass deletes a Class by ID
func (s *ClassService) DeleteClass(id string) error {
	return s.repo.Delete(id)
}
