package models

import (
	"time"
)

type Invoice struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	OrderID         uint      `json:"order_id"`
	TransactionID   string    `json:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	Order       Order       `json:"-"`
	Transaction Transaction `json:"-"`
}