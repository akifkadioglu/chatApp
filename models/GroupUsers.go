package models

import "gorm.io/gorm"

type GroupUser struct {
	gorm.Model
	GroupId int   `json:"group_id"`
	FromId  int   `json:"from_id"`
	IsAdmin bool  `json:"is_admin" gorm:"default:false"`
	Group   Group `json:"group" gorm:"foreignKey:group_id; References:id"`
}
