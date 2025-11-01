package models

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	MovieID uint   `json:"movie_id" gorm:"not null"`
	Movie   Movie  `gorm:"foreignKey:MovieID"`
}
