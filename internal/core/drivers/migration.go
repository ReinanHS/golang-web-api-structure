package drivers

import (
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	err := db.AutoMigrate(
		&entity.User{},
	)

	return err
}
