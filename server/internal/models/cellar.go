package models

import "gorm.io/gorm"

type Cellar struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Location string `json:"location"`
	Wines    []Wine `json:"wines"`
}
