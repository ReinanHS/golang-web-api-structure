package drivers

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"github.com/reinanhs/golang-web-api-structure/internal/core/drivers/mysql"
	"github.com/reinanhs/golang-web-api-structure/internal/core/drivers/sqlite"
	"gorm.io/gorm"
	"strings"
)

type Driver struct {
	Name string
	Db   *gorm.DB
	Ctx  context.Context
}

func New(ctx context.Context) *gorm.DB {
	connectionType := ctx.Value("config").(*config.AppConfig).DBConnection
	connectionType = strings.ToLower(connectionType)

	switch connectionType {
	case "sqlite":
		return sqlite.New(ctx)
	case "mysql":
		return mysql.New(ctx)
	}

	panic("connection type is not supported")
}
