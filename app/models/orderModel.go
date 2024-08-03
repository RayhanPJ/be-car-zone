package models

import (
	"time"
)

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	CarID      uint      `json:"car_id"`
	TotalPrice float64   `json:"total_price"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User User `json:"-"`
	Car  Car  `json:"-"`
}
