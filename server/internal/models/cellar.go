package models

import "gorm.io/gorm"

type Cellar struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Wines     []Wine  `json:"wines"`
}
