package models

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	ClassName string `gorm:"unique;not null"`
}
