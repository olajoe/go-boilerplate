package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Env          string
	Port         string
	Timeout      int
	WriteTimeout int `mapstructure:"WRITE_TIMEOUT"`
	ReadTimeout  int `mapstructure:"READ_TIMEOUT"`
	IdleTimeout  int `mapstructure:"IDLE_TIMEOUT"`
}

func NewConfiguration() *Configuration {
	var cfg Configuration

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.SetDefault("ENV", "local")
	viper.SetDefault("PORT", "3001")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &cfg
}
