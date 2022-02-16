package config

import (
	"os"
	"sync"
	// "github.com/labstack/gommon/log"
	// "github.com/spf13/viper"
)

type AppConfig struct {
	Port     int
	Database struct {
		Driver   string
		Name     string
		Address  string
		Port     int
		Username string
		Password string
	}
	S3Config struct {
		Region     string
		KeyID      string
		AccessKey  string
		BucketName string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8080
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("DB_ADDRESS")
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.S3Config.Region = os.Getenv("S3_REGION")
	defaultConfig.S3Config.KeyID = os.Getenv("S3_KEY_ID")
	defaultConfig.S3Config.AccessKey = os.Getenv("S3_ACCESS_KEY")
	defaultConfig.S3Config.BucketName = os.Getenv("S3_BUCKET_NAME")

	return &defaultConfig
}
