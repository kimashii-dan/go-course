package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title  string  `json:"title" gorm:"not null"`
	Year   int     `json:"year" gorm:"not null"`
	Actors []Actor `json:"actors" gorm:"foreignKey:MovieID"`
}
