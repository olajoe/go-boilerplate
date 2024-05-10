package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Env  string `env:"ENV" envDefault:"local"`
	Port string `env:"PORT" envDefault:"3000"`
}

func NewConfiguration() *Configuration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	c := &Configuration{}

	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	return c
}

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

// func getEnvAsInt(key string, fallback int) int {
// 	valueStr := getEnv(key, "")
// 	if value, err := strconv.Atoi(valueStr); err == nil {
// 		return value
// 	}
// 	return fallback
// }
