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

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config")

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Info("error to load config file, will use default value ", err)
	// 	return &defaultConfig
	// }

	// var finalConfig AppConfig
	// err := viper.Unmarshal(&finalConfig)
	// if err != nil {
	// 	log.Info("failed to extract config, will use default value")
	// 	return &defaultConfig
	// }

	return &defaultConfig
}
