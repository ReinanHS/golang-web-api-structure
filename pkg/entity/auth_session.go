package entity

type AuthSession struct {
	Model
	UserId       uint   `json:"user_id" binding:"required"`
	User         User   `json:"user"`
	Location     string `json:"location" binding:"required"`
	Latitude     string `json:"latitude" binding:"required"`
	Longitude    string `json:"longitude" binding:"required"`
	Device       string `json:"device" binding:"required"`
	UserAgent    string `json:"user_agent" binding:"required"`
	IpAddress    string `json:"ip_address" binding:"required"`
	DeviceId     string `json:"device_id" gorm:"type:varchar;NULL"`
	DeviceIdUuid string `json:"device_id_uuid" gorm:"type:varchar;NULL"`
	IsActive     bool   `json:"is_active" gorm:"default:false"`
}
