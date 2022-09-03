package entity

import "gorm.io/gorm"

type UserCoin struct {
	gorm.Model
	UserId int    `json:"user_id" binding:"required"`
	User   User   `json:"user"`
	Total  int    `gorm:"NOT NULL" json:"total" binding:"required"`
	From   string `gorm:"NOT NULL" json:"from" binding:"required"`
}
