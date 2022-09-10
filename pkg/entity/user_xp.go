package entity

type UserXp struct {
	Model
	UserId uint   `json:"user_id" binding:"required"`
	User   User   `json:"user"`
	Total  int    `gorm:"NOT NULL" json:"total" binding:"required"`
	From   string `gorm:"NOT NULL" json:"from" binding:"required"`
}
