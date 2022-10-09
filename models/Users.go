package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Image    string `json:"image" gorm:"default:null"`
	Name     string `json:"name" gorm:"not null; size:60"`
	Email    string `json:"email" gorm:"not null; size:60; unique" validate:"required,email"`
	Username string `json:"username" gorm:"not null; size:60; unique" validate:"required,min=3"`
	Biography string `json:"biography" gorm:"default:null"`
	Password string `json:"password" gorm:"not null" validate:"required,min=5"`
}
