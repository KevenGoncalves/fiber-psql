package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
)

type EnvVars struct {
	DB_URI      string `mapstructure:"DB_URI"`
	DB_DRIVER   string `mapstructure:"DB_DRIVER"`
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")

	if env == "production" {
		return EnvVars{
			DB_URI:      os.Getenv("DB_URI"),
			DB_DRIVER:   os.Getenv("DB_DRIVER"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if config.DB_URI == "" {
		err = errors.New("DB_URI is required")
		return
	}

	if config.DB_DRIVER == "" {
		err = errors.New("DB_DRIVER is required")
		return
	}

	return
}
