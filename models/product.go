package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:500"`
	Price       float64
}
