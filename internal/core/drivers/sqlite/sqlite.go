package sqlite

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(ctx context.Context) *gorm.DB {
	database := ctx.Value("config").(*config.AppConfig).DBDatabase

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
