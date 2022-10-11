package models

import "gorm.io/gorm"

type GroupMessages struct {
	gorm.Model
	Message string
	GroupId int `json:"group_id"`
	UserId  int `json:"user_id"`
	User    User
}
