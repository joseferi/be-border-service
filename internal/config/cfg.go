package config

import (
	"log"

	"github.com/spf13/viper"
)

func RegisterConfiguration() Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env") // yaml
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("no .env file found or failed to read %v", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Printf("failed unmarshall configurations : %v", err)
	}
	return Cfg
}
