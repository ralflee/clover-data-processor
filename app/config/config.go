package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

//InitAppConfig initialzie app configurations
func InitAppConfig(configPath string) {

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	//parse config file
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Err(err).Msg("config file not found")
	}
}
