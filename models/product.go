package models

import "time"

type Product struct {
	ID           int `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
