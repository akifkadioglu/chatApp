package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null; size:60"`
	Email     string `json:"email" gorm:"not null; size:60"`
	Username  string `json:"username" gorm:"not null; size:60"`
	Biography string `json:"biography"`
	Password  string `json:"password" gorm:"not null;->:false"`
}
