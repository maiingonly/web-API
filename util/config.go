package util

import (
	"github.com/spf13/viper"
)

//Config stores all configuration of the application.
// The values are read by viper from a config file or environment varibles.
type Config struct {
	DBUser        string `mapstructure:"POSTGREST_USER"`
	DBPassword    string `mapstructure:"POSTGREST_PASSWORD"`
	DBName        string `mapstructure:"POSTGREST_DB"`
	DBHost        string `mapstructure:"POSTGREST_HOST"`
	DBPort        string `mapstructure:"POSTGREST_PORT"`
	ServerAddress string `mapstructure:"SERVER_PORT"`
	DBCon         string `mapstructure:"DB"`
}

//LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
