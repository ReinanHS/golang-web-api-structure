package mysql

import (
	"context"
	"fmt"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var lock = &sync.Mutex{}
var (
	mysqlInstance *gorm.DB
)

func connectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbName,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getDriver(ctx context.Context) *gorm.DB {
	if mysqlInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if mysqlInstance == nil {
			log.Println("Creating Mysql single instance now.")

			appConfig := ctx.Value("config").(*config.AppConfig)
			db, err := connectToDB(appConfig.DBUsername, appConfig.DBPassword, appConfig.DBDatabase)

			// unable to connect to database
			if err != nil {
				log.Fatalln(err)
			}

			//defer func(db *sql.DB) {
			//	_ = db.Close()
			//}(db.DB())

			mysqlInstance = db
		} else {
			log.Printf("Single Mysql instance already created")
		}
	} else {
		log.Println("Single Mysql instance already created.")
	}

	return mysqlInstance
}

func New(ctx context.Context) *gorm.DB {
	return getDriver(ctx)
}
