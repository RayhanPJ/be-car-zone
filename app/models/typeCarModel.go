package models

import "time"

type TypeCar struct {
	ID        int
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
