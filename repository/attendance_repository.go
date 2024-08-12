package repository

import (
	"backend_student/models"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(attendance *models.Attendance) error
	FindByID(id uint) (*models.Attendance, error)
	Update(attendance *models.Attendance) error
	Delete(id uint) error
	ReportAttendanceByStudentID(studentID string, startDate, endDate time.Time) ([]models.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) Create(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepository) FindByID(id uint) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Preload("Student").First(&attendance, id).Error
	return &attendance, err
}

func (r *attendanceRepository) Update(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Attendance{}, id).Error
}

func (r *attendanceRepository) ReportAttendanceByStudentID(studentID string, startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Where("student_id = ? AND attendance_date BETWEEN ? AND ?", studentID, startDate, endDate).Find(&attendances).Error
	return attendances, err
}
