package entity

import "gorm.io/gorm"

type AuthAccessToken struct {
	gorm.Model
	AuthSessionId int `json:"auth_session_id" binding:"required"`
	AuthSession   AuthSession
	Scopes        string `json:"scopes" binding:"required" gorm:"type:text"`
	IsRevoked     bool   `json:"is_revoked" binding:"required"`
}
