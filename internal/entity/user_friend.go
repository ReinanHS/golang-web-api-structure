package entity

import "gorm.io/gorm"

type UserFriend struct {
	gorm.Model
	UserId   int  `json:"user_id" binding:"required"`
	User     User `json:"user"`
	FriendId int  `json:"friend_id" binding:"required"`
	Friend   User `json:"friend" gorm:"foreignKey:FriendId"`
}
