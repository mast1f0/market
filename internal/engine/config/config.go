package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	JWT_SECRET  string
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
