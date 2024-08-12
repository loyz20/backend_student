package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Student struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID string    `gorm:"unique;not null"`
	Name      string    `gorm:"not null"`
	Class     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Username  string    `gorm:"null"`
	User      User      `gorm:"foreignKey:Username;references:Username"`
}

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	ClassName string `gorm:"unique;not null"`
}

type StudentClass struct {
	ID        uint    `gorm:"primaryKey"`
	StudentID uint    `gorm:"not null"`
	ClassID   uint    `gorm:"not null"`
	Student   Student `gorm:"foreignKey:StudentID;references:ID"`
	Class     Class   `gorm:"foreignKey:ClassID;references:ID"`
}

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
