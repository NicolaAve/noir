package models

import "gorm.io/gorm"

type Wine struct {
	gorm.Model

	// DATI STRUTTURALI
	Name     string `json:"name" gorm:"not null"`
	CellarID uint   `json:"cellar_id" gorm:"not null"`
	Quantity int    `json:"quantity" gorm:"default:1"`
	Status   string `json:"status" gorm:"type:varchar(20);default:'in_stock'"`

	// DATI IDENTIFICATIVI
	Producer string `json:"producer"`
	Year     int    `json:"year"`
	Type     string `json:"type"`
	Grape    string `json:"grape"`

	// DATI AVANZATI E DEGUSTAZIONE
	CriticScore int    `json:"critic_score"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`

	// Un vino ha piu record di degustazione
	History []TastingHistory `json:"history" gorm:"foreignKey:WineID"`
}
