package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type AppConfig struct {
	AppName  string `mapstructure:"APP_NAME"`
	AppHost  string `mapstructure:"APP_HOST"`
	AppPort  string `mapstructure:"APP_PORT"`
	AppDebug string `mapstructure:"APP_DEBUG"`

	DBDatabase string `mapstructure:"DB_DATABASE"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var lock = &sync.Mutex{}

var (
	configInstance *AppConfig
)

// LoadConfig reads configuration from file or environment variables.
func loadConfig(path string) (config AppConfig, err error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetConfig(ctx context.Context) *AppConfig {

	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			fmt.Println("Creating AppConfig single instance now.")
			configData, err := loadConfig(".")

			if err != nil {
				log.Fatal("cannot load config:", err)
			}

			configInstance = &configData
		} else {
			fmt.Println("Single AppConfig instance already created.")
		}
	} else {
		fmt.Println("Single AppConfig instance already created.")
	}

	return configInstance
}
