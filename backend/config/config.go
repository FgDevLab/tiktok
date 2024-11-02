package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

func loadConfigFile() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}
	})
}

func GetDBConfig() DBConfig {
	loadConfigFile()
	return LoadDBConfig()
}

func GetHTTPConfig() HTTPConfig {
	loadConfigFile()
	return LoadHTTPConfig()
}
