package models

import "gorm.io/gorm"

type UserCountry struct {
	gorm.Model
	CountryId int
	UserId    int
	User      User
	Country   Country
}
