package models

import "time"

type Attendance struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Username       string    `json:"username"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	PhotoURL       string    `json:"photo_url"`
	IsWithinRadius bool      `json:"is_within_radius"`
	CreatedAt      time.Time `json:"created_at"`
}

type MonthlyReport struct {
	Present int `json:"present"`
	Sick    int `json:"sick"`
	Leave   int `json:"leave"`
	Absent  int `json:"absent"`
}
