package models

import "time"

type Animal struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"-"`
	Brinco    string    `gorm:"not null" json:"brinco"`
	BirthDate time.Time `gorm:"not null" json:"birth_date"`
	Sex       string    `gorm:"not null" json:"sex"` // "Fêmea" or "Macho"
	Breed     string    `gorm:"not null" json:"breed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
