package models

import "time"

type Car struct {
	ID          int
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageUrl    string    `json:"image_url"`
	TypeID      uint      `json:"type_id"`
	BrandID     uint      `json:"brand_id"`
	IsSecond    bool      `json:"is_second"`
	Type        TypeCar   `json:"type" gorm:"foreignKey:TypeID"`
	Brand       BrandCar  `json:"brand" gorm:"foreignKey:BrandID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
