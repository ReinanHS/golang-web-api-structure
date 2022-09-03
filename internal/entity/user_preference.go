package entity

import (
	"gorm.io/gorm"
)

type UserPreference struct {
	gorm.Model
	UserId   int    `json:"user_id" binding:"required"`
	User     User   `json:"user"`
	Settings string `gorm:"type:text" json:"settings" binding:"required"`
}
