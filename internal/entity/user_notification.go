package entity

type UserNotification struct {
	Model
	UserId  uint   `json:"user_id" binding:"required"`
	User    User   `json:"user"`
	Payload string `json:"payload" binding:"required"`
	IsRead  bool   `json:"is_read" binding:"required" gorm:"default:false"`
}
