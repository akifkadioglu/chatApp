package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	FromId int `json:"from_id"`
	ToId   int `json:"to_id"`
	From   User `json:"from" gorm:"foreignkey:from_id;references:id;constraint:OnDelete:CASCADE"`
	To     User `json:"to" gorm:"foreignkey:to_id;references:id;constraint:OnDelete:CASCADE"`
}
