package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Image string `json:"image"`
	Name  string `json:"name"`
}
