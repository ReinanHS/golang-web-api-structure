package entity

type AuthAccessToken struct {
	Model
	AuthSessionId uint `json:"auth_session_id" binding:"required"`
	AuthSession   AuthSession
	Scopes        string `json:"scopes" binding:"required" gorm:"type:text"`
	IsRevoked     bool   `json:"is_revoked" binding:"required"`
}
