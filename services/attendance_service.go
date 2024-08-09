package services

import (
	"errors"
	"fmt"
	"time"

	"backend_student/models"
	"backend_student/utils"
)

const TargetLatitude = -7.341582
const TargetLongitude = 108.251175
const TargetRadius = 100.0

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	return utils.Haversine(lat1, lon1, lat2, lon2)
}

func IsWithinRadius(lat, lon float64) bool {
	distance := CalculateDistance(lat, lon, TargetLatitude, TargetLongitude)
	return distance <= TargetRadius
}

func CreateAttendance(username string, latitude, longitude float64, photoURL string) (*models.Attendance, error) {
	isWithinRadius := IsWithinRadius(latitude, longitude)

	if !isWithinRadius {
		return nil, errors.New("out of location bounds")
	}

	attendance := &models.Attendance{
		Username:       username,
		Latitude:       latitude,
		Longitude:      longitude,
		PhotoURL:       photoURL,
		IsWithinRadius: isWithinRadius,
		CreatedAt:      time.Now(),
	}

	// Save attendance to database (replace with actual DB logic)
	// err := db.Create(attendance).Error
	// if err != nil {
	//     return nil, err
	// }

	fmt.Println("Attendance saved:", attendance)
	return attendance, nil
}

func GetMonthlyReport(username string) (*models.MonthlyReport, error) {
	// Replace with actual logic to fetch data from the database
	report := &models.MonthlyReport{
		Present: 20,
		Sick:    5,
		Leave:   3,
		Absent:  2,
	}
	return report, nil
}
