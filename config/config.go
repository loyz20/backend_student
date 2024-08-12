package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Global variables to hold configuration values
	DB        *gorm.DB
	JwtSecret []byte
)

// InitializeConfig initializes configuration values
func InitializeConfig() {
	// Load JWT Secret from environment variables
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	if JwtSecret == nil {
		fmt.Println("JWT_SECRET is not set")
		return
	}

	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	// Load database connection parameters from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	fmt.Println("Database connection established")

	err = DB.AutoMigrate(
	// &models.User{},
	// &models.Class{},
	// &models.Student{},
	// &models.Attendance{},
	// &models.StudentClass{},
	// &models.RefreshToken{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
