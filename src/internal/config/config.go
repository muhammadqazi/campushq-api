package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	ENVIRONMENT   string `mapstructure:"ENVIRONMENT"`
	PORT          string `mapstructure:"PORT"`
	DB_URL        string `mapstructure:"DB_URL"`
	MIGRATION_URL string `mapstructure:"MIGRATION_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (c Config, err error) {

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
