package config

import (
	"github.com/spf13/viper"
)

type HTTPConfig struct {
	Host string
	Port string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host: viper.GetString("http.host"),
		Port: viper.GetString("http.port"),
	}
}
