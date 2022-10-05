package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message string
	FromId  int  `json:"from_id"`
	ToId    int  `json:"to_id"`
	From    User `json:"from" gorm:"foreignKey:from_id; References:id"`
	To      User `json:"to" gorm:"foreignKey:to_id; References:id"`
}
