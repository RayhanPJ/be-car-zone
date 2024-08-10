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
	OrderImage string    `json:"order_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Car        Car       `json:"car" gorm:"foreignKey:CarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type OrderDetail struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	CarID      uint      `json:"car_id"`
	TotalPrice float64   `json:"total_price"`
	Status     bool      `json:"status"`
	OrderImage string    `json:"order_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Car        CarDetail `json:"car"`
}

type CarDetail struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageCar    string    `json:"image_car"`
	Price       float64   `json:"price"`
	TypeID      uint      `json:"type_id"`
	BrandID     uint      `json:"brand_id"`
	IsSecond    bool      `json:"is_second"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
