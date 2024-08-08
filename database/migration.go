package database

import (
	"backend_student/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	log.Println("Database migrated")
}
