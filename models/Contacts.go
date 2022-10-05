package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	From     int  `json:"from"`
	To       int  `json:"to"`
	UserTo   User `json:"userTo" gorm:"foreignkey:to;references:id;constraint:OnDelete:CASCADE"`
	UserFrom User `json:"userFrom" gorm:"foreignkey:from;references:id;constraint:OnDelete:CASCADE"`
}
