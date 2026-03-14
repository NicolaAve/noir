package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string   `json:"first_name" gorm:"not null"`
	LastName  string   `json:"last_name" gorm:"not null"`
	Email     string   `gorm:"unique;not null" json:"email"`
	Password  string   `gorm:"not null" json:"-"`
	Role      string   `gorm:"type:varchar(20);default:'user'" json:"role"`
	Cellars   []Cellar `gorm:"many2many:user_cellars;" json:"cellars"`
}
