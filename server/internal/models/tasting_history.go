package models

import (
	"time"

	"gorm.io/gorm"
)

type TastingHistory struct {
	gorm.Model
	WineID       uint      `json:"wine_id" gorm:"not null"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	ConsumedDate time.Time `json:"consumed_date"`
	Quantity     int       `json:"quantity" gorm:"default:1"`
	Rating       int       `json:"rating"`
	Notes        string    `json:"notes"`
}
