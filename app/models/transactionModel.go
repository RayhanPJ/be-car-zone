package models

import (
	"time"
)

type Transaction struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	OrderID          uint      `json:"order_id"`
	PaymentProvider  string    `json:"payment_provider"`
	TransactionImage string    `json:"transaction_image"`
	NoRek            string    `json:"no_rek"`
	Amount           float64   `json:"amount"`
	TransactionDate  time.Time `json:"transaction_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	Order Order `json:"order" gorm:"foreignKey:OrderID"`
}

type TransactionDetail struct {
	ID               uint        `json:"id"`
	OrderID          uint        `json:"order_id"`
	PaymentProvider  string      `json:"payment_provider"`
	TransactionImage string      `json:"transaction_image"`
	NoRek            string      `json:"no_rek"`
	Amount           float64     `json:"amount"`
	TransactionDate  time.Time   `json:"transaction_date"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Order            OrderDetail `json:"order"`
}
