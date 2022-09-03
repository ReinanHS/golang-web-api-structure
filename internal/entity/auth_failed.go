package entity

import "gorm.io/gorm"

type AuthFailed struct {
	gorm.Model
	UserId    int    `json:"user_id" binding:"required"`
	User      User   `json:"user"`
	Device    string `json:"device" binding:"required"`
	UserAgent string `json:"user_agent" binding:"required"`
	IpAddress string `json:"ip_address" binding:"required"`
}
