package entity

import (
	"time"
)

type User struct {
	Model
	Username        string     `gorm:"type:varchar" json:"username" binding:"required"`
	Name            string     `gorm:"type:varchar" json:"name" binding:"required"`
	Avatar          string     `gorm:"type:varchar" json:"avatar" binding:"required"`
	Bio             string     `gorm:"type:varchar;NULL" json:"bio"`
	Password        string     `gorm:"type:varchar" json:"-" binding:"required"`
	PhoneNumber     string     `gorm:"type:varchar;NULL" json:"phone_number"`
	Email           string     `gorm:"type:varchar;unique_index" json:"email" binding:"required"`
	EmailVerifiedAt *time.Time `gorm:"NULL" json:"email_verified_at"`

	Tracks []*Track `gorm:"many2many:track_user;"`
}
