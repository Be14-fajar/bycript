package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error read config : ", err.Error())
		return nil
	}
	err = viper.Unmarshal(&app)
	if err != nil {
		log.Println("error parse config : ", err.Error())
		return nil
	}
	return &app
}
