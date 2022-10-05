package models

import "gorm.io/gorm"

type Messages struct {
	gorm.Model
	Message  string
	From     int  `json:"from"`
	To       int  `json:"to"`
	UserFrom User `json:"userFrom" gorm:"foreignKey:from; References:id"`
	UserTo   User `json:"userTo" gorm:"foreignKey:to; References:id"`
}
