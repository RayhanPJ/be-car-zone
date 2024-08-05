package models

import "time"

type Car struct {
	ID          int
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	TypeID      int       `json:"type_id"`
	BrandID     int       `json:"brand_id"`
	IsSecond    bool      `json:"is_second"`
	Type        TypeCar   `json:"type" gorm:"foreignKey:TypeID"`
	Brand       BrandCar  `json:"brand" gorm:"foreignKey:BrandID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
