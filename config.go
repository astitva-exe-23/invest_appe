// config.go

package main

import (
	"github.com/spf13/viper"
)

// Config struct holds configuration settings
type Config struct {
	DBUsername      string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
	AlpacaAPIKey    string
	AlpacaSecretKey string
}

func loadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
