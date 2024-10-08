package models

import "time"

type TypeCar struct {
	ID        int
	Name      string    `json:"name"`
	Cars      []Car     `json:"cars,omitempty" gorm:"foreignKey:TypeID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
