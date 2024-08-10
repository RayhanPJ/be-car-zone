package models

import (
	"time"
)

type Invoice struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `json:"order_id"`
	TransactionID uint      `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Order       Order       `json:"order" gorm:"foreignKey:OrderID"`
	Transaction Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
}

type InvoiceDetail struct {
	ID            uint      `json:"id"`
	OrderID       uint      `json:"order_id"`
	TransactionID uint      `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Order       OrderDetail       `json:"order"`
	Transaction TransactionDetail `json:"transaction"`
}
