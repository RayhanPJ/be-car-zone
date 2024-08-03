package models

import "time"

type Car struct {
	ID          int
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	TypeID      uint      `json:"type_id"`
	IsSecond    bool      `json:"is_second"`
	Type        TypeCar   `json:"type" gorm:"foreignKey:TypeID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
