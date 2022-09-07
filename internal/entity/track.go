package entity

type Track struct {
	Model
	Name   string `json:"name" binding:"required"`
	Config string `json:"config" binding:"required" gorm:"type:text;NOT NULL"`

	Users []*User `gorm:"many2many:track_user;"`
}
