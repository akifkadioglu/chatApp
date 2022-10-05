package models

import "gorm.io/gorm"

type UserLanguage struct {
	gorm.Model
	FromId     int       `json:"from_id"`
	LanguageId int       `json:"language_id"`
	User       User      `json:"user" gorm:"foreignKey:from_id; References:id"`
	Language   Language `json:"language" gorm:"foreignKey:language_id; References:id"`
}
