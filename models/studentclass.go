package models

type StudentClass struct {
	ID        uint    `gorm:"primaryKey"`
	StudentID uint    `gorm:"not null"`
	ClassID   uint    `gorm:"not null"`
	Student   Student `gorm:"foreignKey:StudentID;references:ID"`
	Class     Class   `gorm:"foreignKey:ClassID;references:ID"`
}
