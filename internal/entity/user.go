package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(255);NOT NULL;unique_index" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255);NOT NULL" json:"-" binding:"required"`
}

type Users []User
