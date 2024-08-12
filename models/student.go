package models

import "time"

type Student struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID string    `gorm:"unique;not null"`
	Name      string    `gorm:"not null"`
	Class     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Username  string    `gorm:"null"`
	User      User      `gorm:"foreignKey:Username;references:Username"`
}
