package entity

type UserFriend struct {
	Model
	UserId   uint `json:"user_id" binding:"required"`
	User     User `json:"user"`
	FriendId uint `json:"friend_id" binding:"required"`
	Friend   User `json:"friend" gorm:"foreignKey:FriendId"`
}
