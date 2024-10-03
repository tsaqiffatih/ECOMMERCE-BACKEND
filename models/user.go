package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	Email    string `gorm:"unique;size:100"`
	Password string `gorm:"size:100"`
}
