package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type Robot struct {
	ID              uint
	Name            string
	Description     string
	Model           string
	Version         string
	ManufactureCost float64
	ManufactureDate time.Time

	// Các thông số kỹ thuật
	BatteryLife     string
	Width           float64
	Height          float64
	Depth           float64
	Weight          float64
	Material        string
	Color           string
	CPU             string
	RAM             string
	Storage         string
	OperatingSystem string

	// Các thông số mở rộng khác
	SensorCount     int
	CameraCount     int
	HasWiFi         bool
	HasBluetooth    bool
	IsAutonomous    bool
	WaterproofLevel string

	// Thời gian tạo & cập nhật
	CreatedAt time.Time
	UpdatedAt time.Time
}
