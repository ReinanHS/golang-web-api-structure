package entity

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Config string `json:"config" binding:"required" gorm:"type:text;NOT NULL"`

	Users []*User `gorm:"many2many:track_user;"`
}
