package models

import "gorm.io/gorm"

type Wine struct {
	gorm.Model
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Type     string `json:"type"`     // Rosso, Bianco, ecc.
	Producer string `json:"producer"`
	CellarID uint   `json:"cellar_id"`
}