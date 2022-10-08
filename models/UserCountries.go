package models

import "gorm.io/gorm"

type UserCountry struct {
	gorm.Model
	CountryId int     `json:"country_id"`
	FromId    int     `json:"from_id"`
	From      User    `json:"from" gorm:"foreignKey:from_id; References:id"`
	Country   Country `json:"country" gorm:"foreignKey:country_id; References:id"`
}
