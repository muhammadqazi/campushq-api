package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	ENVIRONMENT         string `mapstructure:"ENVIRONMENT"`
	PORT                string `mapstructure:"PORT"`
	DB_URL              string `mapstructure:"DB_URL"`
	MIGRATION_URL       string `mapstructure:"MIGRATION_URL"`
	AUTH0_DOMAIN        string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_CLIENT_ID     string `mapstructure:"AUTH0_CLIENT_ID"`
	AUTH0_CLIENT_SECRET string `mapstructure:"AUTH0_CLIENT_SECRET"`
	AUTH0_CALLBACK_URL  string `mapstructure:"AUTH0_CALLBACK_URL"`
	AUTH0_AUDIENCE      string `mapstructure:"AUTH0_AUDIENCE"`
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
