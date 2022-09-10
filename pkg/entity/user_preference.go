package entity

type UserPreference struct {
	Model
	UserId   uint   `json:"user_id" binding:"required"`
	User     User   `json:"user"`
	Settings string `gorm:"type:text" json:"settings" binding:"required"`
}
