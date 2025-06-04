package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"Database"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"Host"`
	Port     int8   `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Db       string `mapstructure:"DB"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config") // ищет рядом с main.go

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения config.yaml: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Ошибка парсинга config.yaml: %v", err)
	}
}
