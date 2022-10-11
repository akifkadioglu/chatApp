package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	UserId int  `json:"user_id"`
	ToId   int  `json:"to_id"`
	To     User `json:"to" gorm:"foreignkey:to_id;references:id;constraint:OnDelete:CASCADE"`
}
