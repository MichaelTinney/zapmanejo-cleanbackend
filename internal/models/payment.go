// internal/models/payment.go
package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	PayPalOrderID string    `gorm:"unique" json:"paypal_order_id"`
	Amount        float64   `json:"amount"` // in BRL
	Type          string    `json:"type"`   // "monthly" or "lifetime"
	Status        string    `json:"status"` // "COMPLETED", etc.
	CreatedAt     time.Time `json:"created_at"`
}
