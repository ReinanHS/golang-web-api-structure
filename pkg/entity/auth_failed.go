package entity

type AuthFailed struct {
	Model
	UserId    uint   `json:"user_id" binding:"required"`
	User      User   `json:"user"`
	Device    string `json:"device" binding:"required"`
	UserAgent string `json:"user_agent" binding:"required"`
	IpAddress string `json:"ip_address" binding:"required"`
}
