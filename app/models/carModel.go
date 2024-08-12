package models

import "time"

type Car struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageCar    string    `json:"image_car"`
	Price       float64   `json:"price"`
	TypeID      uint      `json:"type_id"`
	BrandID     uint      `json:"brand_id"`
	IsSecond    bool      `json:"is_second"`
	Type        TypeCar   `json:"type" gorm:"foreignKey:TypeID"`
	Brand       BrandCar  `json:"brand" gorm:"foreignKey:BrandID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
