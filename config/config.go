package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(".env")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return v
}
