package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration variables from environment
type Config struct {
	Port string
}

// Load the environment variables and return config struct
func Load() Config {
	viper.AutomaticEnv()

	viper.BindEnv("Port", "SERVER_PORT")
	viper.SetDefault("Port", "8080")

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
