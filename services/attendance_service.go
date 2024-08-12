package services

import (
	"backend_student/models"
	"backend_student/repository"
	"time"
)

type AttendanceService interface {
	CreateAttendance(attendance *models.Attendance) error
	GetAttendanceByID(id uint) (*models.Attendance, error)
	UpdateAttendance(attendance *models.Attendance) error
	DeleteAttendance(id uint) error
	GenerateAttendanceReport(studentID string, startDate, endDate time.Time) (map[string]int, error)
}

type attendanceService struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{attendanceRepo}
}

func (s *attendanceService) CreateAttendance(attendance *models.Attendance) error {
	return s.attendanceRepo.Create(attendance)
}

func (s *attendanceService) GetAttendanceByID(id uint) (*models.Attendance, error) {
	return s.attendanceRepo.FindByID(id)
}

func (s *attendanceService) UpdateAttendance(attendance *models.Attendance) error {
	return s.attendanceRepo.Update(attendance)
}

func (s *attendanceService) DeleteAttendance(id uint) error {
	return s.attendanceRepo.Delete(id)
}

func (s *attendanceService) GenerateAttendanceReport(studentID string, startDate, endDate time.Time) (map[string]int, error) {
	attendances, err := s.attendanceRepo.ReportAttendanceByStudentID(studentID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	report := map[string]int{
		"hadir":             0,
		"sakit":             0,
		"izin":              0,
		"tanpa_keterangan":  0,
		"hari_telat":        0,
		"total_menit_telat": 0,
	}

	for _, attendance := range attendances {
		switch attendance.Status {
		case "hadir":
			report["hadir"]++
			if attendance.LateMinutes > 0 {
				report["hari_telat"]++
				report["total_menit_telat"] += attendance.LateMinutes
			}
		case "sakit":
			report["sakit"]++
		case "izin":
			report["izin"]++
		case "tanpa keterangan":
			report["tanpa_keterangan"]++
		}
	}

	return report, nil
}
