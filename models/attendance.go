package models

import "time"

type Attendance struct {
	ID             uint      `gorm:"primaryKey"`
	StudentID      string    `gorm:"not null"`
	Student        Student   `gorm:"foreignKey:StudentID;references:StudentID"`
	AttendanceDate time.Time `gorm:"not null"`
	AttendanceType string    `gorm:"not null"`  // "masuk" or "pulang"
	Status         string    `gorm:"not null"`  // "hadir", "sakit", "izin", "tanpa keterangan"
	LateMinutes    int       `gorm:"default:0"` // jumlah menit keterlambatan
	Location       string    `gorm:"type:text"`
	SelfieImageUrl string
}
