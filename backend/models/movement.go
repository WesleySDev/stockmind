package models

import "time"

type Movement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
