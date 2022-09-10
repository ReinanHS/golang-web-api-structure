package drivers

import (
	"github.com/reinanhs/golang-web-api-structure/pkg/entity"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	err := db.AutoMigrate(
		&entity.User{},
		&entity.UserPreference{},
		&entity.AuthSession{},
		&entity.AuthAccessToken{},
		&entity.AuthFailed{},
		&entity.UserCoin{},
		&entity.UserXp{},
		&entity.UserFriend{},
		&entity.UserNotification{},
		&entity.Track{},
	)

	return err
}
